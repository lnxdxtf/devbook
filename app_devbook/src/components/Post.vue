<template>
    <div>
        <div class="w-full flex flex-col bg-black bg-opacity-20 rounded-md p-2">
            <div class="w-full flex justify-between items-center text-slate-600 dark:text-slate-300">
                <span>
                    <img :src="`${devbookCDN}/user/profile/imgs/user_${post.author_id}.png`" @error="onImageError"
                        alt="user image" class="w-10 h-10 rounded-full p-2">
                </span>
                <span class="text-[10px]" v-timePassed="post.created_at">
                </span>
            </div>
        </div>
    </div>
</template>
<script lang="ts">
import { Prop, Component, Vue, toNative } from 'vue-facing-decorator';
import { PostDevBookAPI } from '../app/modules/user/interfaces';

@Component({})
class Post extends Vue {
    @Prop
    post!: PostDevBookAPI
    @Prop
    devbookCDN!: string

    onImageError(event: any) {
        event.target.src = `${this.devbookCDN}/user/profile/imgs/user_default_0.png`
    }
}
export default toNative(Post)
</script>
