components {
  id: "enemy_1"
  component: "/game/enemy/enemy_1/enemy_1.script"
}
embedded_components {
  id: "spinemodel"
  type: "spinemodel"
  data: "spine_scene: \"/game/enemy/enemy.spinescene\"\n"
  "default_animation: \"idle\"\n"
  "skin: \"enemy_1\"\n"
  "material: \"/defold-spine/assets/spine.material\"\n"
  ""
}
