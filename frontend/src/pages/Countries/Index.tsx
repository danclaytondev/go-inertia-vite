/* eslint-disable  @typescript-eslint/no-explicit-any */
import Layout from "../../components/Layout";

type Props = {
	countries: string[];
};

export default function Home(props: Props) {
	return (
		<Layout>
			<div className="py-6">
				<ul>
					{props.countries.map((c) => (
						<li>{c}</li>
					))}
				</ul>
			</div>
		</Layout>
	);
}
