/* eslint-disable  @typescript-eslint/no-explicit-any */
import Layout from "../../components/Layout";

type Props = {
	countries: {
		Name: string,
		Flag: string,
	}[]
};

export default function Home(props: Props) {
	return (
		<Layout>
			<div className="py-6 text-lg">
				<ul>
					{props.countries.map(c => (
						<li key={c.Name}>{c.Flag} {c.Name}</li>
					))}
				</ul>
			</div>
		</Layout>
	);
}
