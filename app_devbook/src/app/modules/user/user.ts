import { Method, RequestOptions, ResponseJsonSuccess, fetchAPIDevBook } from "../../utils/fetch";
import { NotificationError, NotificationSuccess } from "../../utils/notification";
import { LoginFormDevBookAPI, RandomUserDevBookAPI, RegisterFormDevBookAPI, UserDevBook, UserDevBookAPI } from "./interfaces";
/**
 * User class.
 * It implements the UserDevBookAPI interface.
 * User class is used to create to do login and populate the user data if login is successful.
 * This class provides methods to login, logout, password recovery, etc. All About this user.
 * 
 * @implements UserDevBookAPI
 * @class User
 * @module User
 * 
 * @example
 * const user = new User()
 * 
 */



/**
 * WHAT TO DO HERE?
 * REDUCE THE CODE
 * LOGIN FUNCTION POPULATE REDUCE TO A SINGLE FUNCTION THAT POPULATE THE USER DATA
 */

export default class User implements UserDevBook {
    id?: number;
    name?: string;
    nick?: string;
    email?: string;
    token?: string;
    token_exp?: number;
    created_at?: string | Date;

    async login(form: LoginFormDevBookAPI): Promise<boolean> {
        try {
            const response = await fetchAPIDevBook({ method: Method.POST, path: '/login', body: form } as RequestOptions) as ResponseJsonSuccess
            let user_builder: UserDevBookAPI = {
                id: response.data.id,
                token: response.data.token,
                token_exp: response.data.exp
            }
            if (!user_builder.id) {
                throw new Error('Could not login, please try again later')
            }

            const headers = { 'Authorization': `Bearer ${user_builder.token}` }
            const userDataResponse = await fetchAPIDevBook({ method: Method.GET, path: `/users/${user_builder.id}`, headers: headers } as RequestOptions) as ResponseJsonSuccess

            if (!userDataResponse.data.email) {
                throw new Error('Could not login, please try again later')
            }
            user_builder.email = userDataResponse.data.email
            user_builder.name = userDataResponse.data.name
            user_builder.nick = userDataResponse.data.nick
            user_builder.created_at = userDataResponse.data.created_at
            this.setUser(user_builder)
            this.rememberLogin()
            NotificationSuccess('Login successful')
            setInterval(async () => await this.checkToken(user_builder.token_exp!), 10000)
            return true

        } catch (error) {
            console.log(error)
            NotificationError('Error attempting to login, please try again later')
            return false
        }

    }


    private async checkToken(exp: number): Promise<void> {
        if (exp < Date.now()) {
            await this.logout()
        }
    }

    private setUser(user: UserDevBookAPI): void {
        this.id = user.id
        this.name = user.name
        this.nick = user.nick
        this.email = user.email
        this.token = user.token
        this.token_exp = user.token_exp
        this.created_at = user.created_at
    }

    public getLoggedUser(): boolean {
        if (!localStorage.getItem('user')) { return false }
        if (!localStorage.getItem('token-exp')) { return false }
        if (!localStorage.getItem('token')) { return false }
        const exp = parseInt(localStorage.getItem('token-exp') as string)
        if (exp < Date.now()) { return false }
        const user = JSON.parse(localStorage.getItem('user') as string) as UserDevBookAPI
        this.id = user.id
        this.name = user.name
        this.nick = user.nick
        this.email = user.email
        this.token = user.token
        this.created_at = user.created_at
        return true
    }

    private rememberLogin(): void {
        const user = {
            id: this.id,
            name: this.name,
            nick: this.nick,
            email: this.email,
            token: this.token,
            token_exp: this.token_exp,
            created_at: this.created_at
        }
        localStorage.setItem('token', user.token as string)
        localStorage.setItem('token-exp', user.token_exp! as unknown as string)
        localStorage.setItem('user', JSON.stringify(user))
    }

    async getUserPosts(): Promise<void> {
        try {
            const headers = { 'Authorization': `Bearer ${this.token}` }
            const response = await fetchAPIDevBook({ method: Method.GET, path: '/posts', headers } as RequestOptions) as ResponseJsonSuccess
            return response.data
        } catch (error) {
            console.log(error)
            NotificationError('Error attempting to get feed, please try again later')
            return []
        }
    }

    public async logout(): Promise<void> {
        delete this.id
        delete this.name
        delete this.email
        delete this.token
        delete this.token_exp
        delete this.created_at
        localStorage.removeItem('user')
        localStorage.removeItem('token')
        localStorage.removeItem('token-exp')
        window.location.reload()
    }

    async register(form: RegisterFormDevBookAPI): Promise<boolean> {
        try {
            const response = await fetchAPIDevBook({ method: Method.POST, path: '/users', body: form } as RequestOptions) as ResponseJsonSuccess
            if (response.data) {
                return true
            }
            NotificationError("Could not register, please try again later")
            return false

        } catch (err) {
            console.log(err)
            NotificationError("Could not register, please try again later")
            return false
        }
    }

    async followUser(id:number): Promise<void> {
        const headers = { 'Authorization': `Bearer ${this.token}` }
        await fetchAPIDevBook({ method: Method.POST, path: `/users/${id}/follow`, headers } as RequestOptions)
    }

}