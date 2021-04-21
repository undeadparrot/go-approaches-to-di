from pyramid.config import Configurator
from pyramid.response import Response

def get_cats(request):
    return Response(json={
        "cats": [
            "Puss puss",
            "Mew meow"
        ]
    })

def make_app():
    config = Configurator()
    config.add_route('root', '/cats')
    config.add_view(get_cats, route_name='root')
    return config.make_wsgi_app()
