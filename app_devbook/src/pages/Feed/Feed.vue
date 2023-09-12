<template>
    <div>
        <div class="w-full h-full flex flex-col gap-4 p-4">
            <Post v-for="post in feed" :key="post.id" :post="post" />
        </div>
    </div>
</template>
<script lang="ts">
import { Component, Vue, toNative } from 'vue-facing-decorator';
import FeedMixin from './feedMixin'
import Post from '../../components/Post.vue';
@Component({
    mixins: [FeedMixin],
    components: { Post }
})
class Feed extends Vue { }
export default toNative(Feed)
</script>