import { GetServerSideProps } from "next";
import axios from "axios";

const OrdersPage = (props: any) => {
    console.log(props.orders)
    return (
        <div>
            Listagens de orders...{props.name}
        </div>
    );
};

export default OrdersPage;

export const getServerSideProps: GetServerSideProps = async (context) => {

    const {data} = await axios.get('http//localhost:3000/orders');

    return {
        props: {
            orders: data
        }
    }
}