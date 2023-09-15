import { Vue, Component, toNative } from 'vue-facing-decorator';
import store from '../../app/store/store';

@Component
class FeedMixin extends Vue {

    get feed() {
        return store.state.user.feed
    }
    get devbookCDN() {
        return store.state.devbookCDN
    }
    async beforeMount() {
        await store.dispatch('user/GetFeed')
    }

}

export default toNative(FeedMixin)