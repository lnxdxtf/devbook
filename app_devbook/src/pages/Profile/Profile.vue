<template lang>
    <div>
        <div class="w-full h-full p-4 bg-green-500"></div>
    </div>
</template>
<script lang="ts">
import { Component, Vue, toNative } from 'vue-facing-decorator';
@Component
class Profile extends Vue { }
export default toNative(Profile)
</script>