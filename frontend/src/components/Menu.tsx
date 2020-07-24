import React, { Component } from 'react';
import axios from 'axios';

interface State {
    isLoggedIn: boolean
    user: {
        id: number,
        name: string
    }
}

class Menu extends Component<{}, State> {
    state: State = {
        isLoggedIn: false,
        user: {
            id: 0,
            name: "test"
        }
    };

    // GET / をしてログインしているならユーザー情報を取得する
    componentDidMount() {
        const params = {
            email: "a@a.com",
            password: "password"
        };

        axios.get('http://localhost:8000/', {headers: {'Authorization': ''}})
            .then((response) => {
                const userData = response.data.User;
                if(userData != null){
                    this.setState({
                        user: {
                            id: userData.id,
                            name: userData.name
                        }
                    })
                }
            })
            .catch(() => {
                console.log("index fail");
            });
    }

    render() {
        const isLoggedIn = this.state.isLoggedIn;
        let msg: string;
        if (this.state.user.id == 0) {
            msg = "ログインしてません。";
        } else{
            msg = "ログイン済みです。";
        }

        return (
            <div className="Menu">
                {msg}
            </div>
        );
    }
}

export default Menu;

