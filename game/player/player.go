components {
  id: "player"
  component: "/game/player/player.script"
}
embedded_components {
  id: "co"
  type: "collisionobject"
  data: "type: COLLISION_OBJECT_TYPE_KINEMATIC\n"
  "mass: 0.0\n"
  "friction: 0.1\n"
  "restitution: 0.5\n"
  "group: \"player\"\n"
  "mask: \"enemies\"\n"
  "mask: \"ground\"\n"
  "embedded_collision_shape {\n"
  "  shapes {\n"
  "    shape_type: TYPE_BOX\n"
  "    position {\n"
  "      x: -4.0\n"
  "      y: 58.0\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 0\n"
  "    count: 3\n"
  "  }\n"
  "  data: 31.257132\n"
  "  data: 54.027588\n"
  "  data: 10.0\n"
  "}\n"
  ""
}
embedded_components {
  id: "co_radar"
  type: "collisionobject"
  data: "type: COLLISION_OBJECT_TYPE_TRIGGER\n"
  "mass: 0.0\n"
  "friction: 0.1\n"
  "restitution: 0.5\n"
  "group: \"radar\"\n"
  "mask: \"enemies\"\n"
  "embedded_collision_shape {\n"
  "  shapes {\n"
  "    shape_type: TYPE_SPHERE\n"
  "    position {\n"
  "      x: -3.0\n"
  "      y: 61.0\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 0\n"
  "    count: 1\n"
  "  }\n"
  "  data: 72.45213\n"
  "}\n"
  ""
}
embedded_components {
  id: "spinemodel"
  type: "spinemodel"
  data: "spine_scene: \"/game/player/player.spinescene\"\n"
  "default_animation: \"run\"\n"
  "skin: \"body/body_1\"\n"
  "material: \"/defold-spine/assets/spine.material\"\n"
  ""
  position {
    x: -7.0
    y: 11.0
  }
  scale {
    z: -1.0
  }
}
