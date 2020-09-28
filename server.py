from flask import Flask, render_template_string

app = Flask(__nam__)

@app.route("/")
def index():
    return render_template_string("<h1> hello, world </h1>")
