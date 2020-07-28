import React, { Component } from 'react';
import { Link, RouteComponentProps } from "react-router-dom";
import axios from "axios";

type MyPageProps = {} & RouteComponentProps<{id: string}>;

interface MyPageState {
    books: [{
        ID: number,
        Sentence: string,
        Title: string,
        Author: string,
        Pages: number,
        Harts: number,
        UpdatedAd: any
    }]
}

// 自分の情報を表示するマイページ
class MyPage extends Component<MyPageProps, MyPageState> {
    state: MyPageState = {
        books: [{
            ID: 0,
            Sentence: "",
            Title: "",
            Author: "",
            Pages: 0,
            Harts: 0,
            UpdatedAd: ""
        }]
    };

    constructor(props: MyPageProps) {
        super(props);
    }

    componentDidMount() {
        axios.get('http://localhost:8000/users/books', {params: {id: this.props.match.params.id}})
            .then((response) => {
                const books: any[] = response.data.Books;
                if(books != null){
                    books.forEach((key, index) => {
                       console.log(books[index]);
                    });
                }
            })
            .catch(() => {
                console.log("books get fail");
            });

    }

    render() {
        return (
            <div>
                マイページです。
                <Link to="/">ホームへ</Link>
                {(() => {
                    if (this.state.books) {
                        return (
                            <div>
                            </div>
                        )
                    }
                })()}
            </div>
        );
    }
}

export default MyPage;
