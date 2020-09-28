from flask import Flask, render_template_string, render_template


def create_app():

    app = Flask(__name__)

    @app.route("/")
    def index():
        return render_template("index.html")

    @app.route("/article")
    def article():
        return render_template_string("<h1> hello, article </h1>")

    @app.route("/article/<article_id>")
    def show_article(article_id):
        return render_template_string("<h1> hello, article nr: %s </h1>" % article_id)

    return app
