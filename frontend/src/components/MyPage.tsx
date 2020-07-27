import React, { Component } from 'react';
import { Link, RouteComponentProps } from "react-router-dom";

type MyPageProps = {} & RouteComponentProps<{id: string}>;

// 自分の情報を表示するマイページ
class MyPage extends Component<MyPageProps, {}> {
    constructor(props: MyPageProps) {
        super(props);
    }

    componentDidMount() {
        console.log(this.props.match.params.id);
        // console.log(this.props);
    }

    render() {
        return (
            <div>
                マイページです。
                <Link to="/">ホームへ</Link>
            </div>
        );
    }
}

export default MyPage;
