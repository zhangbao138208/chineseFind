<template>
  <div>
    <div label="我爱你">我爱你</div><div label="随风倒十分">随风倒十分</div>
    <div label="我爱你">我爱你</div><div title="随风倒十分">随风倒十分</div>
<!--    <div label="我爱你">我爱你</div><div title="随风倒十分">随风倒十分</div>-->
    <div label="我爱你">我爱你</div><div :title="'随风倒十分'">随风倒十分</div>
    <div label="我爱你">我爱你</div><div title='随风倒十分'>{{'随风倒十分'}}</div>
    <div label="我爱你">我爱你</div><div label="随风倒十分">随风倒十分</div>
    <div label="我爱你">我爱你</div><div placeholder="随风倒十分">随风倒十分</div>
  </div>
</template>
<script>
//我爱你
function f() {
  alert("我爱你")
}
</script>