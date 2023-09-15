import { Vue, Component, toNative } from 'vue-facing-decorator';
import store from '../../app/store/store';
import { LoginFormDevBookAPI } from '../../app/modules/user/interfaces';
import { NotificationError } from '../../app/utils/notification';

/**
 * This function is used to login a user.
 * The login methods in this function call the login action in the store.
 */
@Component
class LoginMixin extends Vue {
    public email: string = ''
    public pswrd: string = ''

    public async login() {
        if (this.validate()) {
            const login_ok = await store.dispatch('user/LoginAction', { email: this.email, pswrd: this.pswrd } as LoginFormDevBookAPI)
            if (login_ok) {
                return this.$router.push('/')
            }
        }
        this.pswrd = ''
        return
    }

    private validate(): boolean {
        if (this.email.length === 0 || this.pswrd.length === 0) {
            NotificationError('Email and password are required, please fill them')
            return false
        }
        return true
    }

}

export default toNative(LoginMixin)