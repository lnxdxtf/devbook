import { Vue, Component, toNative } from 'vue-facing-decorator';
import store from '../../app/store/store';
import { NotificationError } from '../../app/utils/notification';
import { RegisterFormDevBookAPI } from '../../app/modules/user/interfaces';
/**
 * This function is used to register a new user.
 * The register methods in this function call the register action in the store.
 */
@Component
class RegisterMixin extends Vue {

    public email: string = ''
    public pswrd: string = ''
    public pswrd_confirm: string = ''
    public name: string = ''
    public nick: string = ''
    public image: any

    public async register() {
        if (this.validate()) {
            const register_ok = await store.dispatch('user/RegisterAction', {
                email: this.email,
                pswrd: this.pswrd,
                pswrd_confirm: this.pswrd_confirm,
                name: this.name,
                nick: this.nick,
                image: this.image
            } as RegisterFormDevBookAPI)
            if (register_ok) {
                return this.$router.push('/login')
            }
        }
        this.reset()
        return

    }

    triggerInput() {
        document.getElementById('register-profile-image-input')?.click()
    }

    private watchImage() {
        document.getElementById('register-profile-image-input')?.addEventListener('change', (e: any) => {
            const file = e.target.files[0]
            if (file) {
                const reader = new FileReader()
                reader.onload = (e) => {
                    document.getElementById('register-profile-image')!.src = e.target.result
                    this.image = e.target.result
                }
                reader.readAsDataURL(file)
            }
        })
    }

    mounted() {
        this.watchImage()
    }

    private validate(): boolean {
        if (this.email.length === 0 || this.pswrd.length === 0
            || this.pswrd_confirm.length === 0 || this.name.length === 0
            || this.nick.length === 0) {
            NotificationError('Email, password, password confirm, name and nick are required, please fill them')
            this.reset()
            return false
        }
        if (this.pswrd !== this.pswrd_confirm) {
            NotificationError('Password and password confirm must be the same')
            return false
        }
        return true
    }

    private reset() {
        this.email = ''
        this.pswrd = ''
        this.pswrd_confirm = ''
        this.name = ''
        this.nick = ''
        this.image = null
    }
}

export default toNative(RegisterMixin)