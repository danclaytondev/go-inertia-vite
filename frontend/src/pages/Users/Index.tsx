/* eslint-disable  @typescript-eslint/no-explicit-any */
import Layout from "../../components/Layout";

type Props = {
	users: string[];
};

export default function Home(props: Props) {
	return (
		<Layout>
			<div className="py-6">
				<ul>
					{props.users.map((u) => (
						<li>{u}</li>
					))}
				</ul>
			</div>
		</Layout>
	);
}
