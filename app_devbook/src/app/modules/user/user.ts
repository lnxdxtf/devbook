import { Method, RequestOptions, fetchAPIDevBook } from "../../utils/fetch";
import { LoginFormDevBookAPI, UserDevBook } from "./interfaces";
/**
 * User class.
 * It implements the UserDevBookAPI interface.
 * User class is used to create to do login and populate the user data if login is successful.
 * Every time someone enters the application, a new instance of this class is created.
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
export default class User implements UserDevBook {
    id?: number;
    name?: string;
    email?: string;
    token?: string;
    created_at?: string | Date;


    async login(form: LoginFormDevBookAPI) {
        try {
            const response = await fetchAPIDevBook({ method: Method.POST, path: '/login', body: form } as RequestOptions)
            if ('data' in response) {
                console.log(response)
            } else {
                console.log(response);
            }
        } catch (error) {
            console.log(error)
        }


    }

    public async logout() {
        delete this.id
        delete this.name
        delete this.email
        delete this.token
        delete this.created_at
    }

}