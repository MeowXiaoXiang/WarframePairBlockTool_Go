<template>
  <div class="content text-center">
    <div class="mb-2">
      <label>Warframe 配對限制器（強制主機）</label>
      <br>

      <label> 封鎖 UDP 輸出 Port </label>
    </div>

    <div class="mb-2">
      <label class="p-2" :class="[status?'statusTrueLabel':'statusFalseLabel']"> 目前狀態 :
        {{ status ? '配對目前正常' : '配對已阻斷' }}</label>
    </div>

    <div class="mb-2">
      <label> 選擇您Warframe內的UDP埠 </label>
      <br>

      <select v-model="port_select">
        <option v-for="(v,k) in port_combo" :value="k">{{ v }}</option>
      </select>
      <br>

      <label>（若是預設可以不用更動）</label>
    </div>

    <div class="buttons d-flex flex-column px-2">
      <button class="btn btn-secondary" @click="checkFireWall">查看防火牆狀態</button>
      <button class="btn btn-danger" @click="disableMate">阻止配對[封鎖UDP輸出]</button>
      <button class="btn btn-success" @click="enableMate">恢復配對[恢復UDP輸出]</button>
    </div>
  </div>
</template>

<script setup>
import {ref} from "vue";
import {CheckRule, CreateRule, DelRule, GetRuleStatus} from "../wailsjs/go/main/App.js";

const status = ref(true)
const port_select = ref(0)
const port_combo = ["4950 & 4955",
  "4960 & 4965",
  "4970 & 4975",
  "4980 & 4985",
  "4990 & 4995",
  "3074 & 3080"]

function disableMate() {
  CreateRule(port_select.value).then((e) => {
    status.value = e
  })
}

function checkFireWall() {
  CheckRule()
}

function enableMate() {
  DelRule().then((e) => {
    status.value = e
  })
}

GetRuleStatus().then((e) => {
  status.value = e
})
</script>

<style scoped>
.content {

}

.statusTrueLabel {
  background-color: rgb(0, 205, 102);
}

.statusFalseLabel {
  background-color: rgb(238, 44, 44);
}

.buttons {
  width: 18em;
  margin: auto;
}
</style>
