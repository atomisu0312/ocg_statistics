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
def _get_current_id_from_store():
    """
    SSM Parameter Storeから値を取得する
    """
    ssm_path = os.environ.get('OCG_CURRENT_ID_SSM_PATH', 'invalid-ssm-path')
    ssm_hook = SsmHook(aws_conn_id='aws_default', region_name='ap-northeast-1')
    parameter_value = ssm_hook.get_parameter_value(parameter=ssm_path)
    print(f"Got parameter value: {parameter_value}")
    return parameter_value

@task(task_id="get_max_id_from_store")
def _get_max_id_from_store():
    """
    SSM Parameter Storeから値を取得する
    """
    ssm_path = os.environ.get('OCG_MAX_ID_SSM_PATH', 'invalid-ssm-path')
    ssm_hook = SsmHook(aws_conn_id='aws_default', region_name='ap-northeast-1')
    parameter_value = ssm_hook.get_parameter_value(parameter=ssm_path)
    print(f"Got parameter value: {parameter_value}")
    return parameter_value

@task(task_id="get_delta_id_from_store")
def _get_delta_id_from_store():
    """
    SSM Parameter Storeから値を取得する
    """
    ssm_path = os.environ.get('OCG_DELTA_ID_SSM_PATH', 'invalid-ssm-path')
    ssm_hook = SsmHook(aws_conn_id='aws_default', region_name='ap-northeast-1')
    parameter_value = ssm_hook.get_parameter_value(parameter=ssm_path)
    print(f"Got parameter value: {parameter_value}")
    return parameter_value

@task(task_id="update_current_id")
def _update_current_id(current_id_str: str, delta_id_str: str):
    """
    XComからcurrent_idを取得し、インクリメントしてSSM Parameter Storeに書き戻す
    """
    print(f"Current ID from XCom: {current_id_str}")

    if current_id_str is None:
        print("Could not find current_id in XCom. Skipping update.")
        return

    try:
        current_id = int(current_id_str)
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
def _alarm_excess_id():
    print("Alarm excess id")

@task.branch(task_id="choose_branch_path")
def _choose_branch_path(current_id: str, delta_id: str, max_id: str):
    """
    This function determines which task to execute next based on a random choice.
    It returns the task_id of the chosen branch.
    """
    
    print(f"Current ID from XCom: {current_id}")
    
    print(f"Delta ID from XCom: {delta_id}")
    
    print(f"Max ID from XCom: {max_id}")
    
    if int(current_id) + int(delta_id) >= int(max_id):
        return 'alarm_excess_id'
    else:
        return 'invoke_lambda_function' 

with DAG(
    dag_id="ocg_db_insert",
    start_date=airflow.utils.dates.days_ago(0),
    schedule="0 5 * * *",
) as dag:
    start = EmptyOperator(task_id="start")

    start_fetch_param = EmptyOperator(task_id="start_fetch_param")

    end = EmptyOperator(task_id="end")

    # デコレータで定義したタスクを呼び出し
    get_current_id_task = _get_current_id_from_store()
    get_max_id_task = _get_max_id_from_store()
    get_delta_id_task = _get_delta_id_from_store()
    
    choose_branch_path = _choose_branch_path(get_current_id_task, get_delta_id_task, get_max_id_task)
    alarm_excess_id_task = _alarm_excess_id()
    
    # 個別のタスクを定義
    invoke_lambda_function_task = LambdaInvokeFunctionOperator(
        task_id="invoke_lambda_function",
        region_name='ap-northeast-1',
        function_name=os.environ.get('OCG_LAMBDA_FUNCTION_NAME', 'invalid-lambda-function-name'),
        payload='''{
            "startId": {{ task_instance.xcom_pull(task_ids='get_current_id_from_store') }},
            "delta":{{ task_instance.xcom_pull(task_ids='get_delta_id_from_store') }}
        }''',
    )
    
    update_id_task = _update_current_id(get_current_id_task, get_delta_id_task)
    
    end_update_task = EmptyOperator(task_id="end_update_task", trigger_rule="one_success")

    start >> start_fetch_param >> [get_current_id_task, get_max_id_task, get_delta_id_task]
    choose_branch_path >> [invoke_lambda_function_task, alarm_excess_id_task]
    invoke_lambda_function_task >> update_id_task
    
    [update_id_task, alarm_excess_id_task] >> end_update_task
    end_update_task >> end
    
    
# DAGオブジェクトを明示的に作成
dag = dag