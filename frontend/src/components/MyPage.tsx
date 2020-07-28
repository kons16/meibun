import React, { Component } from 'react';
import { Link, RouteComponentProps } from "react-router-dom";
import axios from "axios";

type MyPageProps = {} & RouteComponentProps<{id: string}>;

interface MyPageState {
    books: any
}

// 自分の情報を表示するマイページ
class MyPage extends Component<MyPageProps, MyPageState> {
    constructor(props: MyPageProps) {
        super(props);
        this.state = {
            books: null
        };
    }

    componentDidMount() {
        axios.get('http://localhost:8000/users/books', {params: {id: this.props.match.params.id}})
            .then((response) => {
                const books: any[] = response.data.Books;
                if(books != null){
                    this.setState({
                        books: books
                    })
                }
                console.log(this.state.books);
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
                        const bookItems: any = [];
                        this.state.books.forEach((key: any, index: number) => {
                            bookItems.push(this.state.books[index].Sentence)
                        });

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
