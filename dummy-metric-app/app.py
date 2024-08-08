
from flask import Flask, Response
from prometheus_client import Gauge, generate_latest
import psutil
import random

app = Flask(__name__)

# Define Prometheus metrics
cpu_usage_gauge = Gauge('cpu_usage', 'CPU usage percentage')
memory_usage_gauge = Gauge('memory_usage', 'Memory usage percentage')
latency_gauge = Gauge('latency', 'Network Latency')


def generate_synthetic_latency(min_latency=10, max_latency=300):
    """Generate a synthetic latency value in milliseconds."""
    return random.uniform(min_latency, max_latency)


@app.route('/metrics')
def metrics():
    # Update metrics
    cpu_usage_gauge.set(psutil.cpu_percent())
    memory_usage_gauge.set(psutil.virtual_memory().percent)

    # Generate a synthetic latency value
    synthetic_latency = generate_synthetic_latency()
    latency_gauge.set(synthetic_latency)

    # Generate latest metrics
    return Response(generate_latest(), mimetype='text/plain')

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5000)