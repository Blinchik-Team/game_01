components {
  id: "cocon"
  component: "/game/bullets/cocon.script"
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"yellow\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/main/main.atlas\"\n"
  "}\n"
  ""
  scale {
    x: 0.5
    y: 0.5
  }
}
embedded_components {
  id: "co"
  type: "collisionobject"
  data: "type: COLLISION_OBJECT_TYPE_KINEMATIC\n"
  "mass: 0.0\n"
  "friction: 0.1\n"
  "restitution: 0.5\n"
  "group: \"player_bullet\"\n"
  "mask: \"enemies\"\n"
  "mask: \"ground\"\n"
  "embedded_collision_shape {\n"
  "  shapes {\n"
  "    shape_type: TYPE_BOX\n"
  "    position {\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 0\n"
  "    count: 3\n"
  "  }\n"
  "  data: 5.8322716\n"
  "  data: 6.214213\n"
  "  data: 10.0\n"
  "}\n"
  ""
}
