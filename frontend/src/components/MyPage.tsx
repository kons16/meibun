import React, { Component } from 'react';
import { Link, RouteComponentProps } from "react-router-dom";
import axios from "axios";
import Book from './Book';

type MyPageProps = {} & RouteComponentProps<{id: string}>;

interface MyPageState {
    books: any
    myID: number
    urlID: number
}

// 自分の情報を表示するマイページ
class MyPage extends Component<MyPageProps, MyPageState> {
    constructor(props: MyPageProps) {
        super(props);
        this.state = {
            books: [],
            myID: 0,
            urlID: 0
        };
    }

    componentDidMount() {
        // ユーザーに基づくbookの取得
        axios.get('http://localhost:8000/users/books', {params: {id: this.props.match.params.id}})
            .then((response) => {
                const books: any[] = response.data.Books;
                if(books != null){
                    this.setState({
                        books: books,
                    })
                }
                console.log(this.state.books);
            })
            .catch(() => {
                console.log("books get fail");
            });

        // マイページに表示するログインしているユーザー情報の取得
        axios.get('http://localhost:8000/', {withCredentials: true})
            .then((response) => {
                const userData = response.data.User;
                if(userData != null){
                    this.setState({
                        myID: userData.ID,
                        urlID: parseInt(this.props.match.params.id)
                    })
                }
            })
            .catch(() => {
                console.log("index fail");
            });
    }

    // Bookコンポーネントに各名文情報を渡いて表示
    render() {
        return (
            <div>
                {this.state.urlID === this.state.myID &&
                    <div>マイページです</div>
                }

                <Link to="/">ホームへ</Link>
                {(() => {
                    const bookItems: any = [];
                    // console.log(this.state.books);
                    this.state.books.forEach((key: any, index: number) => {
                        bookItems.push(
                            <Book
                                key={index}
                                id={index}
                                sentence={this.state.books[index].Sentence}
                                title={this.state.books[index].Title}
                                author={this.state.books[index].Author}
                                pages={this.state.books[index].Pages}
                                harts={this.state.books[index].Harts}
                            />
                        )
                    });

                    if (this.state.books) {
                        return (
                            <div>
                                {bookItems}
                            </div>
                        )
                    }
                })()}
            </div>
        );
    }
}

export default MyPage;
