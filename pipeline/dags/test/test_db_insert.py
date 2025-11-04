import datetime
import os
import pytest
from unittest.mock import patch

from airflow.models.dag import DAG
from airflow.decorators import task

# The tasks to be tested
from ocg.db_insert import _choose_branch_path, _update_current_id

@pytest.fixture
def dag_factory():
    """Provides a factory to create clean, unique DAGs for each test."""
    def _factory(dag_id: str):
        return DAG(
            dag_id,
            default_args={"start_date": datetime.datetime(2023, 1, 1)},
            schedule=None,
            catchup=False,
        )
    return _factory

# --- Tests for _choose_branch_path ---

def test_choose_branch_path_invoke_lambda(dag_factory):
    """Tests that the 'invoke_lambda_function' branch is chosen correctly."""
    dag = dag_factory("test_branch_invoke")

    # 1. Dummy upstream tasks to provide XComs
    @task(task_id="get_current_id_from_store", dag=dag)
    def dummy_get_current_id(**context):
        context["task_instance"].xcom_push(key="current_id", value="100")

    @task(task_id="get_delta_id_from_store", dag=dag)
    def dummy_get_delta_id(**context):
        context["task_instance"].xcom_push(key="delta_id", value="10")

    @task(task_id="get_max_id_from_store", dag=dag)
    def dummy_get_max_id(**context):
        context["task_instance"].xcom_push(key="max_id", value="200")

    # 2. The task to test
    xcom_arg = _choose_branch_path(dag=dag)
    task_to_test = xcom_arg.operator

    # 3. Dummy downstream tasks for the branch operator to find
    @task(task_id="invoke_lambda_function", dag=dag)
    def dummy_invoke_lambda():
        pass

    @task(task_id="alarm_excess_id", dag=dag)
    def dummy_alarm_excess():
        pass

    # 4. Set all dependencies
    [dummy_get_current_id().operator, dummy_get_delta_id().operator, dummy_get_max_id().operator] >> task_to_test
    task_to_test >> [dummy_invoke_lambda().operator, dummy_alarm_excess().operator]

    # 5. Run the test
    dag.test()

    # 6. Assert the state of the downstream tasks
    dagrun = dag.get_last_dagrun()
    ti_invoke = dagrun.get_task_instance("invoke_lambda_function")
    ti_alarm = dagrun.get_task_instance("alarm_excess_id")

    assert ti_invoke.state == "success"
    assert ti_alarm.state == "skipped"

def test_choose_branch_path_alarm_excess_id(dag_factory):
    """Tests that the 'alarm_excess_id' branch is chosen correctly."""
    dag = dag_factory("test_branch_alarm")

    @task(task_id="get_current_id_from_store", dag=dag)
    def dummy_get_current_id(**context):
        context["task_instance"].xcom_push(key="current_id", value="195")

    @task(task_id="get_delta_id_from_store", dag=dag)
    def dummy_get_delta_id(**context):
        context["task_instance"].xcom_push(key="delta_id", value="10")

    @task(task_id="get_max_id_from_store", dag=dag)
    def dummy_get_max_id(**context):
        context["task_instance"].xcom_push(key="max_id", value="200")

    xcom_arg = _choose_branch_path(dag=dag)
    task_to_test = xcom_arg.operator

    @task(task_id="invoke_lambda_function", dag=dag)
    def dummy_invoke_lambda():
        pass

    @task(task_id="alarm_excess_id", dag=dag)
    def dummy_alarm_excess():
        pass

    [dummy_get_current_id().operator, dummy_get_delta_id().operator, dummy_get_max_id().operator] >> task_to_test
    task_to_test >> [dummy_invoke_lambda().operator, dummy_alarm_excess().operator]

    dag.test()

    dagrun = dag.get_last_dagrun()
    ti_invoke = dagrun.get_task_instance("invoke_lambda_function")
    ti_alarm = dagrun.get_task_instance("alarm_excess_id")

    assert ti_invoke.state == "skipped"
    assert ti_alarm.state == "success"

# --- Tests for _update_current_id (passing tests are left as is) ---

@patch('boto3.client')
def test_update_current_id_success(mock_boto3_client, dag_factory):
    """Tests successful execution of _update_current_id."""
    mock_ssm_client = mock_boto3_client.return_value
    dag = dag_factory("test_update_success")
    
    @task(task_id="get_current_id_from_store", dag=dag)
    def dummy_get_current_id(**context):
        context["task_instance"].xcom_push(key="current_id", value="100")

    @task(task_id="get_delta_id_from_store", dag=dag)
    def dummy_get_delta_id(**context):
        context["task_instance"].xcom_push(key="delta_id", value="10")

    xcom_arg = _update_current_id(dag=dag)
    task_to_test = xcom_arg.operator
    
    [dummy_get_current_id().operator, dummy_get_delta_id().operator] >> task_to_test

    with patch.dict('os.environ', {'OCG_CURRENT_ID_SSM_PATH': 'test_path'}):
        dag.test()

    mock_ssm_client.put_parameter.assert_called_once_with(
        Name='test_path',
        Value='110',
        Type='String',
        Overwrite=True
    )

@patch('boto3.client')
def test_update_current_id_no_delta_id(mock_boto3_client, dag_factory):
    """Tests _update_current_id with a default delta of 1."""
    mock_ssm_client = mock_boto3_client.return_value
    dag = dag_factory("test_update_no_delta")

    @task(task_id="get_current_id_from_store", dag=dag)
    def dummy_get_current_id(**context):
        context["task_instance"].xcom_push(key="current_id", value="100")

    @task(task_id="get_delta_id_from_store", dag=dag)
    def dummy_get_delta_id(**context):
        context["task_instance"].xcom_push(key="delta_id", value=None)

    xcom_arg = _update_current_id(dag=dag)
    task_to_test = xcom_arg.operator

    [dummy_get_current_id().operator, dummy_get_delta_id().operator] >> task_to_test

    with patch.dict('os.environ', {'OCG_CURRENT_ID_SSM_PATH': 'test_path'}):
        dag.test()

    mock_ssm_client.put_parameter.assert_called_once_with(
        Name='test_path',
        Value='101',
        Type='String',
        Overwrite=True
    )
