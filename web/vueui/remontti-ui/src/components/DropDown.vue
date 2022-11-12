<template>
  <div class="dropdown">
    <DropDownLine
        v-for="item in items"
        v-bind:title="item.title"
        v-bind:link="item.link"
    />
  </div>
</template>

<script>

import DropDownLine from "./DropDownLine.vue"

export default {
  name: "DropDown",
  components: {DropDownLine},
  props: ['items'],
  methods: {}
}
</script>

<style scoped>
.dropdown {
  display: flex;
  flex-direction: column;
  background: var(--drop-down-background);
  color: var(--drop-down-text-color);
  font-size: 18px;
  position: absolute;
  top: 115%;
  right: 0px;
  width: max-content;
  /*border-bottom: 1px solid red;*/
  /*border-left: 1px solid red;*/
  /*border-right: 1px solid red;*/
  /*border-top: 1px solid red;*/
  /*-webkit-box-shadow: 0px 0px 12px 5px rgba(34, 60, 80, 0.2);*/
  /*-moz-box-shadow: 0px 0px 12px 5px rgba(34, 60, 80, 0.2);*/
  box-shadow: var(--box-shadow-offset-0)
}

/*.dropdown::after {*/
/*  content: "";*/
/*  background: inherit;*/
/*  position: absolute;*/
/*  width: 15px;*/
/*  height: 15px;*/
/*  transform: rotate(45deg);*/
/*  right: 3px;*/
/*  top: -7px;*/
/*  z-index: -1;*/
/*  border: 1px solid red;*/
/*}*/

/*.dropdown::before {*/
/*  content: "";*/
/*  top: -1px;*/
/*  left: -1px;*/
/*  z-index: -1;*/
/*  height: 100%;*/
/*  position: absolute;*/
/*  width:100%;*/
/*  border:1px solid red;*/
/*}*/

</style>