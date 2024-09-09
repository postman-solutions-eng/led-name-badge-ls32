from lednamebadge import * 

creator = SimpleTextAndIcons()
scene_x_bitmap = creator.bitmap("Hallo :HEART2: World!")
scene_y_bitmap = creator.bitmap("Complete example ahead.")
#your_own_stuff = create_own_bitmap_data()

lengths = (scene_x_bitmap[1], scene_y_bitmap[1])
buf = array('B')
buf.extend(LedNameBadge.header(lengths, (3,), (0,), (0,1,0), (0,0,1), 100))
buf.extend(scene_x_bitmap[0])
buf.extend(scene_y_bitmap[0])
#buf.extend(your_own_stuff.bytes)
LedNameBadge.write(buf)