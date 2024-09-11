from flask import Flask, request
from lednamebadge import SimpleTextAndIcons, LedNameBadge
from array import array

app = Flask(__name__)

@app.route('/display-text', methods=['POST'])
def display_text():
    data = request.get_json()
    text = data.get('text', '')

    creator = SimpleTextAndIcons()
    scene_bitmap = creator.bitmap(text)

    buf = array('B')
    buf.extend(LedNameBadge.header([scene_bitmap[1]], [4], [0], [0], [0], 100))
    buf.extend(scene_bitmap[0])
    LedNameBadge.write(buf)

    return {'status': 'Text displayed on LED'}, 200

@app.route('/predefined-icons', methods=['GET'])
def get_predefined_icons():
    creator = SimpleTextAndIcons()
    icons = list(creator.bitmap_named.keys())
    return {'icons': icons}, 200

if __name__ == '__main__':
    app.run(port=5000)
