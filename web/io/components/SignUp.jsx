import React from 'react'
import ReactDOM from 'react-dom'

import UsersActions from '../actions/users'
import Row from '../../core/components/Row.jsx'
import Columns from '../../core/components/Columns.jsx'

export default class SignUp extends React.Component {
    constructor() {
        super()
        this._signUp = this._signUp.bind(this)
    }

    render() {
        return (
            <div className="signup-content">
                <Row>
                    <Columns className="small-offset-1 small-5 description">
                        <h2 className="title">Create a new account</h2>
                        <p>Space is an user management microservice. We aim to provide a secure and reliable authentication and authorization system.</p>
                        <p>By clicking "Sign Up", you agree to our <a href="//quatrolabs.com/terms-of-service">terms of service</a> and <a href="//quatrolabs.com/privacy-policy">privacy policy</a>. We will send you account related emails occasionally.</p>
                    </Columns>
                    <Columns className="small-5 end">
                        <form className="form-sign-up" action="." method="post" ref="form" onSubmit={this._signUp}>
                            {
                                this.props.validationFailed ? (
                                    <p className="error-message">Validation failed</p>
                                ) : null
                            }
                            <Row>
                                <Columns className="small-6">
                                    <input type="text" name="first_name" placeholder="First Name" />
                                </Columns>
                                <Columns className="small-6">
                                    <input type="text" name="last_name" placeholder="Last Name" />
                                </Columns>
                            </Row>
                            <input type="hidden" name="action_token" value="" />
                            <input type="text" name="username" placeholder="Username" />
                            <input type="email" name="email" placeholder="Email" />
                            <input type="password" name="password" placeholder="Password" />
                            <button type="submit" className="button expand" onClick={this._signUp}>Sign Up</button>
                        </form>
                    </Columns>
                </Row>
            </div>
        );
    }

    _signUp(e) {
        e.preventDefault()
        UsersActions.signUp(new FormData(ReactDOM.findDOMNode(this.refs.form)))
    }
}
