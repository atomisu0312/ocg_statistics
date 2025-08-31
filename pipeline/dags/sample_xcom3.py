import uuid

import airflow

from airflow import DAG
from airflow.operators.empty import EmptyOperator
from airflow.operators.python import PythonOperator
import logging

logger = logging.getLogger(__name__)

def _train_model(**context):
    model_id = str(uuid.uuid4())
    context["task_instance"].xcom_push(key="model_id", value=model_id)

    logger.info(f"Training model {model_id}")
    
def _deploy_model(**context):
    model_id = context["task_instance"].xcom_pull(
        task_ids="train_model", key="model_id"
    )
    logger.info(f"Deploying model {model_id}")

with DAG(
    dag_id="10_xcoms_3",
    start_date=airflow.utils.dates.days_ago(3),
    schedule_interval="@daily",
) as dag:
    start = EmptyOperator(task_id="start")

    fetch_sales = EmptyOperator(task_id="fetch_sales")
    clean_sales = EmptyOperator(task_id="clean_sales")

    fetch_weather = EmptyOperator(task_id="fetch_weather")
    clean_weather = EmptyOperator(task_id="clean_weather")

    join_datasets = EmptyOperator(task_id="join_datasets")

    train_model = PythonOperator(task_id="train_model", python_callable=_train_model)

    deploy_model = PythonOperator(task_id="deploy_model", python_callable=_deploy_model)

    start >> [fetch_sales, fetch_weather]
    fetch_sales >> clean_sales
    fetch_weather >> clean_weather
    [clean_sales, clean_weather] >> join_datasets
    join_datasets >> train_model >> deploy_model

# DAGオブジェクトを明示的に作成
dag = dag
