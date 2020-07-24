import React, { Component } from 'react';
import axios from 'axios';

interface State {
    isLoggedIn: boolean
}

class Menu extends Component<{}, State> {
    state: State = {
        isLoggedIn: false
    };

    componentDidMount() {
        const params = {
            email: "a@a.com",
            password: "password"
        };

        axios({
            method: 'POST',
            url: 'http://localhost:8000/signin',
            data: params
        })
            .then(function (response) {
                console.log(response);
            })
            .catch(function (error) {
                console.log(error);
            });
    }

    render() {
        const isLoggedIn = this.state.isLoggedIn;
        let msg: string;
        if (isLoggedIn) {
            msg = "ログイン済みです";
        } else{
            msg = "ログインしてません";
        }

        return (
            <div className="Menu">
                {msg}
            </div>
        );
    }
}

export default Menu;

