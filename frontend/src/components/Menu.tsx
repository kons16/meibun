import React, { Component } from 'react';
import axios from 'axios';

interface State {
    isLoggedIn: boolean
    user: {
        id: number,
        name: string,
        isLoggedIn: boolean
    }
}

class Menu extends Component<{}, State> {
    state: State = {
        isLoggedIn: false,
        user: {
            id: 0,
            name: "test",
            isLoggedIn: false
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
                            name: userData.name,
                            isLoggedIn: true
                        }
                    })
                }
            })
            .catch(() => {
                console.log("index fail");
            });
    }

    render() {
        return (
            <div className="Menu">
                Topページ
            </div>
        );
    }
}

export default Menu;
