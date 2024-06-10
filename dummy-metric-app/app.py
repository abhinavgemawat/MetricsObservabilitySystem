
from flask import Flask, Response
from prometheus_client import Gauge, generate_latest
import psutil

app = Flask(__name__)

# Define Prometheus metrics
cpu_usage_gauge = Gauge('cpu_usage', 'CPU usage percentage')
memory_usage_gauge = Gauge('memory_usage', 'Memory usage percentage')

@app.route('/metrics')
def metrics():
    # Update metrics
    cpu_usage_gauge.set(psutil.cpu_percent())
    memory_usage_gauge.set(psutil.virtual_memory().percent)

    # Generate latest metrics
    return Response(generate_latest(), mimetype='text/plain')

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5000)