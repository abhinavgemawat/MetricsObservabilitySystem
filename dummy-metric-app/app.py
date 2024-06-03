from flask import Flask, jsonify, request
from prometheus_client import start_http_server, Counter, Gauge, Summary
import random
import time
import threading

app = Flask(__name__)

# Create metrics
REQUEST_COUNT = Counter('request_count', 'Total request count')
RESPONSE_TIME = Summary('response_time_seconds', 'Response time in seconds')
CPU_USAGE = Gauge('cpu_usage_percentage', 'CPU usage in percentage')
MEMORY_USAGE = Gauge('memory_usage_percentage', 'Memory usage in percentage')

@app.route('/')
def home():
    return "Dummy Metric App"

@app.route('/metrics')
def metrics():
    return jsonify({
        "request_count": REQUEST_COUNT._value.get(),
        "cpu_usage_percentage": CPU_USAGE._value.get(),
        "memory_usage_percentage": MEMORY_USAGE._value.get()
    })

@app.route('/simulate', methods=['GET'])
@RESPONSE_TIME.time()
def simulate():
    REQUEST_COUNT.inc()
    cpu_usage = random.uniform(0, 100)
    memory_usage = random.uniform(0, 100)
    CPU_USAGE.set(cpu_usage)
    MEMORY_USAGE.set(memory_usage)
    time.sleep(random.uniform(0.1, 0.5))  # Simulate some processing time
    return jsonify({
        "cpu_usage": cpu_usage,
        "memory_usage": memory_usage
    })

def generate_random_metrics():
    end_time = time.time() + 10 * 60  # 10 minutes from now
    while time.time() < end_time:
        cpu_usage = random.uniform(0, 100)
        memory_usage = random.uniform(0, 100)
        CPU_USAGE.set(cpu_usage)
        MEMORY_USAGE.set(memory_usage)
        REQUEST_COUNT.inc()
        time.sleep(1)  # Update metrics every second

if __name__ == '__main__':
    # Start Prometheus metrics server on a new port
    start_http_server(9301)
    
    # Start a thread to generate random metrics
    metrics_thread = threading.Thread(target=generate_random_metrics)
    metrics_thread.start()
    
    # Start Flask server on a new port
    app.run(host="0.0.0.0", port=5101, debug=False)
