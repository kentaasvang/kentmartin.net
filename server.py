from flask import Flask, render_template_string


def create_app():

    app = Flask(__name__)

    @app.route("/")
    def index():
        return render_template_string("<h1> hello, world </h1>")

    return app
