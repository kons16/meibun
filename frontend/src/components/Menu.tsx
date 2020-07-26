import React, { Component } from 'react';
import axios from 'axios';
import { Link } from "react-router-dom";

interface State {
    isLoggedIn: boolean
    user: {
        id: number,
        name: string,
    }
}

class Menu extends Component<{}, State> {
    state: State = {
        isLoggedIn: false,
        user: {
            id: 0,
            name: "test",
        }
    };

    // GET / をしてログインしているならユーザー情報を取得する
    componentDidMount() {
        axios.get('http://localhost:8000/check_user', {withCredentials: true})
            .then((response) => {
                const userData = response.data.User;
                if(userData != null){
                    this.setState({
                        isLoggedIn: true,
                        user: {
                            id: userData.ID,
                            name: userData.Name,
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
                Topページ<br/>
                {this.state.isLoggedIn &&
                    <span>{this.state.user.name}</span>
                }
                <Link to="/login">ログイン</Link> <br/>
                <Link to="/signup">新規登録</Link>
            </div>
        );
    }
}

export default Menu;
