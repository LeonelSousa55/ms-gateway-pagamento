import { GetServerSideProps } from "next";
import axios from "axios";
import { Button } from "@mui/material";

const OrdersPage = (props: any) => {
    console.log(props.orders)
    return (
        <div>
            Listagens de orders...{props.name}
            <Button variant="contained">Contained</Button>
        </div>
    );
};

export default OrdersPage;

export const getServerSideProps: GetServerSideProps = async (context) => {

    const { data } = await axios.get('http://localhost:3000/orders', {
        headers: {
            'x-token': 'wm9jhrn14l'
        }
    });

    return {
        props: {
            orders: data
        }
    }
}