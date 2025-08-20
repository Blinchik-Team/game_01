components {
  id: "enemy"
  component: "/game/enemy/enemy.script"
}
embedded_components {
  id: "co"
  type: "collisionobject"
  data: "type: COLLISION_OBJECT_TYPE_TRIGGER\n"
  "mass: 0.0\n"
  "friction: 0.1\n"
  "restitution: 0.5\n"
  "group: \"enemies\"\n"
  "mask: \"player\"\n"
  "mask: \"radar\"\n"
  "mask: \"player_bullet\"\n"
  "mask: \"ground\"\n"
  "embedded_collision_shape {\n"
  "  shapes {\n"
  "    shape_type: TYPE_BOX\n"
  "    position {\n"
  "      x: -2.0\n"
  "      y: 58.0\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 0\n"
  "    count: 3\n"
  "  }\n"
  "  data: 29.744211\n"
  "  data: 58.56132\n"
  "  data: 10.0\n"
  "}\n"
  ""
}
embedded_components {
  id: "spinemodel"
  type: "spinemodel"
  data: "spine_scene: \"/game/enemy/enemy.spinescene\"\n"
  "default_animation: \"idle\"\n"
  "skin: \"\"\n"
  "material: \"/defold-spine/assets/spine.material\"\n"
  ""
  position {
    x: -7.0
    y: 11.0
  }
}
