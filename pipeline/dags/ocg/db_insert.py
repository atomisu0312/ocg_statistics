import os
import uuid
import json

import airflow

from airflow import DAG
from airflow.operators.empty import EmptyOperator
from airflow.operators.python import PythonOperator
from airflow.providers.amazon.aws.hooks.ssm import SsmHook
from airflow.providers.amazon.aws.operators.lambda_function import LambdaInvokeFunctionOperator
from airflow.decorators import task, task_group

@task(task_id="get_current_id_from_store")
def _get_current_id_from_store(**context):
    """
    SSM Parameter Storeから値を取得する
    """
    ssm_path = os.environ.get('OCG_CURRENT_ID_SSM_PATH', 'invalid-ssm-path')
    ssm_hook = SsmHook(aws_conn_id='aws_default', region_name='ap-northeast-1')
    parameter_value = ssm_hook.get_parameter_value(parameter=ssm_path)
    context["task_instance"].xcom_push(key="current_id", value=parameter_value)
    print(f"Got parameter value: {parameter_value}")
    return parameter_value

@task(task_id="get_max_id_from_store")
def _get_max_id_from_store(**context):
    """
    SSM Parameter Storeから値を取得する
    """
    ssm_path = os.environ.get('OCG_MAX_ID_SSM_PATH', 'invalid-ssm-path')
    ssm_hook = SsmHook(aws_conn_id='aws_default', region_name='ap-northeast-1')
    parameter_value = ssm_hook.get_parameter_value(parameter=ssm_path)
    context["task_instance"].xcom_push(key="max_id", value=parameter_value)
    print(f"Got parameter value: {parameter_value}")
    return parameter_value

@task(task_id="get_delta_id_from_store")
def _get_delta_id_from_store(**context):
    """
    SSM Parameter Storeから値を取得する
    """
    ssm_path = os.environ.get('OCG_DELTA_ID_SSM_PATH', 'invalid-ssm-path')
    ssm_hook = SsmHook(aws_conn_id='aws_default', region_name='ap-northeast-1')
    parameter_value = ssm_hook.get_parameter_value(parameter=ssm_path)
    context["task_instance"].xcom_push(key="delta_id", value=parameter_value)
    print(f"Got parameter value: {parameter_value}")
    return parameter_value

@task(task_id="update_current_id")
def _update_current_id(**context):
    """
    XComからcurrent_idを取得し、インクリメントしてSSM Parameter Storeに書き戻す
    """
    current_id_str = context["task_instance"].xcom_pull(
        task_ids="get_current_id_from_store", key="current_id"
    )
    print(f"Current ID from XCom: {current_id_str}")

    if current_id_str is None:
        print("Could not find current_id in XCom. Skipping update.")
        return

    try:
        current_id = int(current_id_str)
        delta_id_str = context["task_instance"].xcom_pull(
            task_ids="get_delta_id_from_store", key="delta_id"
        )
        print(f"Delta ID from XCom: {delta_id_str}")
        
        if delta_id_str is None:
            print("Could not find delta_id in XCom. Using default value 1.")
            delta_id = 1
        else:
            delta_id = int(delta_id_str)
            
        new_id = current_id + delta_id
        print(f"Current ID: {current_id}, Delta ID: {delta_id}, New ID: {new_id}")

        import boto3
        ssm_path = os.environ.get('OCG_CURRENT_ID_SSM_PATH', 'invalid-ssm-path')
        ssm_client = boto3.client('ssm', region_name='ap-northeast-1')
        ssm_client.put_parameter(
            Name=ssm_path,
            Value=str(new_id),
            Type='String',
            Overwrite=True
        )
        print(f"Successfully updated parameter {ssm_path} to {new_id}")

    except (ValueError, TypeError) as e:
        print(f"Error converting current_id to integer: {e}")
        raise

@task(task_id="alarm_excess_id")
def _alarm_excess_id(**context):
    print("Alarm excess id")

@task.branch(task_id="choose_branch_path")
def _choose_branch_path(**kwargs):
    """
    This function determines which task to execute next based on a random choice.
    It returns the task_id of the chosen branch.
    """
    
    current_id = kwargs["task_instance"].xcom_pull(
        task_ids="get_current_id_from_store", key="current_id"
    )
    print(f"Current ID from XCom: {current_id}")
    
    delta_id = kwargs["task_instance"].xcom_pull(
        task_ids="get_delta_id_from_store", key="delta_id"
    )
    print(f"Delta ID from XCom: {delta_id}")
    
    max_id = kwargs["task_instance"].xcom_pull(
        task_ids="get_max_id_from_store", key="max_id"
    )
    print(f"Max ID from XCom: {max_id}")
    
    if int(current_id) + int(delta_id) >= int(max_id):
        return 'alarm_excess_id'
    else:
        return 'invoke_lambda_function' 

with DAG(
    dag_id="ocg_db_insert",
    start_date=airflow.utils.dates.days_ago(0),
    schedule="*/10 * * * *",
) as dag:
    start = EmptyOperator(task_id="start")

    start_fetch_param = EmptyOperator(task_id="start_fetch_param")
    
    end_fetch_param = EmptyOperator(task_id="end_fetch_param")

    end = EmptyOperator(task_id="end")

    # デコレータで定義したタスクを呼び出し
    get_current_id_task = _get_current_id_from_store()
    get_max_id_task = _get_max_id_from_store()
    get_delta_id_task = _get_delta_id_from_store()
    
    choose_branch_path = _choose_branch_path()
    alarm_excess_id_task = _alarm_excess_id()
    
    # 個別のタスクを定義
    invoke_lambda_function_task = LambdaInvokeFunctionOperator(
        task_id="invoke_lambda_function",
        region_name='ap-northeast-1',
        function_name=os.environ.get('OCG_LAMBDA_FUNCTION_NAME', 'invalid-lambda-function-name'),
        payload='''{
            "startId": {{ task_instance.xcom_pull(task_ids='get_current_id_from_store', key='current_id') }},
            "delta":{{ task_instance.xcom_pull(task_ids='get_delta_id_from_store', key='delta_id') }}
        }''',
    )
    
    update_id_task = _update_current_id()
    
    # 依存関係を設定
    invoke_lambda_function_task >> update_id_task

    start >> start_fetch_param >> [get_current_id_task, get_max_id_task, get_delta_id_task] >> end_fetch_param
    end_fetch_param >> choose_branch_path >> [invoke_lambda_function_task, alarm_excess_id_task]
    update_id_task >> end
    alarm_excess_id_task >> end
    
# DAGオブジェクトを明示的に作成
dag = dag