from diagrams import Cluster, Diagram
from diagrams.gcp.analytics import BigQuery
from diagrams.gcp.compute import Run
from diagrams.gcp.network import LoadBalancing
from diagrams.gcp.database import Spanner, Memorystore
from diagrams.gcp.operations import Monitoring

with Diagram("", show=False):

    lb = LoadBalancing("Google Cloud Load Balancing")

    with Cluster("Application"):
        run = Run("game-api")
        spanner = Spanner("game")
    
    lb >> run
    run >> spanner

    with Cluster("Data"):
        bq = BigQuery("BigQuery")
        monitoring = Monitoring("Cloud Logging")
        monitoring >> bq
    run >> monitoring
