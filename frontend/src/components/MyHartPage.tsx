import React, { Component } from 'react';
import axios from "axios";

// 自分はハートしたbook一覧を表示するページ
class MyHartPage extends Component<{}, {}> {

    componentDidMount() {
        axios.get('http://localhost:8000/get_my_harts', {withCredentials: true})
            .then((response) => {
                const myHartBooks: any[] = response.data.myHartBooks;
                if(myHartBooks != null){
                    console.log(myHartBooks);
                }
                // console.log(this.state.books);
            })
            .catch(() => {
                console.log("books get fail");
            });
    }

    render() {
        return (
            <div>
            </div>
        );
    }
}

export default MyHartPage;
