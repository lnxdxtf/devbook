<template>
    <div>

        <div class="w-full flex flex-col bg-black bg-opacity-20 rounded-md p-2 gap-4">
            <div class="w-full flex justify-between items-center text-slate-600 dark:text-slate-300">
                <span class="cursor-pointer flex items-center gap-2">
                    <img :id="`user_${post.author_id}`" alt="user image" class="w-10 h-10 rounded-full p-2">
                    <span>{{ post.author_nick }}</span>
                    <span class="font-bold text-teal-400">USER</span>
                </span>
                <span class="text-[10px]" v-timePassed="post.created_at">
                </span>
            </div>
            <div class="w-full flex justify-center items-center">
                <img :id="`post_${post.id}`" alt="post image" class="rounded-md w-1/5">
            </div>
            <div class="w-full flex items-center gap-4">

            </div>
            <div class="w-full px-2 text-slate-800 dark:text-slate-200">{{ post.title }}</div>
            <div class="w-full p-2 text-slate-800 dark:text-slate-200">{{ post.content }}</div>
            <div class="w-full flex gap-4 p-4 border-t-[.5px] border-gray-700 dark:text-white">
                <span :class="{ 'text-devbook-main-one': post.likes }"
                    class="p-1 cursor-pointer text-xl flex flex-col gap-1 items-center">
                    <i class="fa-solid fa-mug-saucer"></i>
                    <span class="text-xs">{{ post.likes ?? 0 }}</span>
                </span>
                <span class="p-1 cursor-pointer text-xl flex flex-col gap-1 items-center">
                    <i class="fa-solid fa-code"></i>
                    <span class="text-xs">{{ 0 }}</span>
                </span>
                <span class="p-1 cursor-pointer text-xl flex flex-col gap-1 items-center">
                    <i class="fa-solid fa-code-fork"></i>
                    <span class="text-xs">{{ 0 }}</span>
                </span>
            </div>
            
        </div>

    </div>
</template>
<script lang="ts">
import { Prop, Component, Vue, toNative } from 'vue-facing-decorator';
import { PostDevBookAPI } from '../app/modules/user/interfaces';
import { ImageDevBookClass, setImgUrl } from '../app/utils/image';
import Loader from './Loader.vue';
@Component({
    components: { Loader }
})
class Post extends Vue {
    @Prop
    post!: PostDevBookAPI
    @Prop
    devbookCDN!: string


    async setImg(elementId: string, imgType: ImageDevBookClass, imgID: number) {
        await setImgUrl(elementId, imgType, imgID)
    }

    async mounted() {
        await setImgUrl(`user_${this.post.author_id}`, ImageDevBookClass.UserProfile, this.post.author_id)
        await setImgUrl(`post_${this.post.id}`, ImageDevBookClass.Post, this.post.id)
    }

}
export default toNative(Post)
</script>
