import { Vue, Component, toNative } from 'vue-facing-decorator';
import store from '../../app/store/store';
import { LoginFormDevBookAPI } from '../../app/modules/user/interfaces';
import { NotificationError } from '../../app/utils/notification';

@Component
class LoginMixin extends Vue {
    public email: string = ''
    public pswrd: string = ''

    public async login() {
        if (this.email.length === 0 || this.pswrd.length === 0) {
            NotificationError('Email and password are required, please fill them')
            return
        }
        await store.dispatch('user/LoginAction', { email: this.email, pswrd: this.pswrd } as LoginFormDevBookAPI)
        this.$router.push('/')
    }

}

export default toNative(LoginMixin)