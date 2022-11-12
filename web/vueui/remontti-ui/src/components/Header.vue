<script>

import LeftHeader from "./LeftHeader.vue"
import CenterHeader from "./CenterHeader.vue";
import RightHeader from "./RightHeader.vue";
import {computed} from "vue";

export default {
  name: "Header",
  components: {RightHeader, CenterHeader, LeftHeader},
  props : {
    translations : Object
  },
  data() {
    return {
      translations1 : computed( () => this.translations)
    }
  }
}
</script>

<template>
  <div class="header">
    <LeftHeader/>
    <CenterHeader/>
    <RightHeader v-bind:translations="translations1"/>
  </div>
</template>

<style scoped>

.header {
  border-bottom: 2px solid var(--light-gray);
  min-height: 50px;
  display:flex;
  justify-content: space-between;
  align-items: center;
}

.header-part {
  min-width:100px;
}

</style>