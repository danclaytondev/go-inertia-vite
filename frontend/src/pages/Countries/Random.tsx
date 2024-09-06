/* eslint-disable  @typescript-eslint/no-explicit-any */
import Layout from "../../components/Layout";
import { router } from "@inertiajs/react";

type Props = {
	countries: {
		Name: string;
		Flag: string;
	}[];
};

export default function Home(props: Props) {
	return (
		<Layout>
			<div className="py-6 text-lg">
				<ul>
					{props.countries.map((c) => (
						<li key={c.Name}>
							{c.Flag} {c.Name}
						</li>
					))}
				</ul>
				<div>
					<button
						className="block bg-purple-600 text-white text-sm font-medium rounded-sm px-4 py-2 mt-2"
						onClick={() => router.reload()}
					>
						Refresh
					</button>
				</div>
			</div>
		</Layout>
	);
}
