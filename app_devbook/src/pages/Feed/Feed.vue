<template>
    <div>
        <div class="w-full h-full flex flex-col gap-8 p-4">
            <template v-if="feed">
                <Post v-for="post in feed" :key="post.id" :post="post" :devbookCDN="devbookCDN" />
            </template>
            <template v-else>
                <div class="w-full h-full flex flex-col justify-center items-center text-black dark:text-white">
                    <span class="text-2xl font-bold items-center">
                        <i class="fa-regular fa-folder-open"></i> No Posts
                    </span>
                    <span class="text-xl items-center">
                        Follow Someone to see their posts
                    </span>
                </div>
            </template>
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