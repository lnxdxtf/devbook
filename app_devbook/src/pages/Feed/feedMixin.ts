import { Vue, Component, toNative } from 'vue-facing-decorator';
import store from '../../app/store/store';
import { NotificationError } from '../../app/utils/notification';
import { PostDevBookAPI } from '../../app/modules/user/interfaces';

@Component
class FeedMixin extends Vue {

    get feed() {
        return store.state.user.feed
    }
    async beforeMount() {
        await store.dispatch('user/GetFeed')
    }

}

export default toNative(FeedMixin)