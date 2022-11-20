<template>
  <div class="icon-with-dropdown-block">
    <div class="icon-with-dropdown">
      <img v-on:click.stop="dropdownToggle" :src="image(src)"/>
      <DropDown v-show="show" :items="items"/>
    </div>
  </div>
</template>

<script>
//TODO: Заменить иконки на svg чтоб можно было кинуть тень при наведении
import DropDown from "./DropDown.vue"

export default {
  data() {
    return {
      show: false,
    }
  },
  name: "IconDropdown",
  props: ['src', 'items'],
  components: {DropDown},
  methods: {
    image(img) {
      return new URL(img, import.meta.url).href
    },
    dropdownToggle() {

      this.show = !this.show

      document.addEventListener("click", this.dropdownClose)
      document.addEventListener("keyup", this.dropdownClose)

    },
    dropdownClose(event) {

      if (event.type === "keyup" && event.key !== 'Escape') {
        return
      }

      this.show = false

      document.removeEventListener("click", this.dropdownClose)
      document.removeEventListener("keyup", this.dropdownClose)

    }
  }
}
</script>

<style scoped>
img {
  width: 25px;
  height: 25px;
  cursor: pointer;
}

img:hover {
  transform: scale(1.2);
}

.icon-with-dropdown {
  position: relative;
  display: flex;
}

.icon-with-dropdown-block {
  display: flex;
}
</style>