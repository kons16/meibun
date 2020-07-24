import React, { Component } from 'react';

interface State {
    isLoggedIn: boolean
}

class Menu extends Component<{}, State> {
    state: State = {
        isLoggedIn: false
    };

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

