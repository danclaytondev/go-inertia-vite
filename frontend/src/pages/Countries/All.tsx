/* eslint-disable  @typescript-eslint/no-explicit-any */
import Layout from "../../components/Layout";
import { useForm } from "@inertiajs/react";

type Props = {
	countries: {
		Name: string;
		Flag: string;
	}[];
};

export default function Home(props: Props) {
	const { data, setData, post, processing } = useForm({
		name: "",
		code: "",
	});

	function submit(e: any) {
		e.preventDefault();
		post("/countries");
	}

	return (
		<Layout>
			<div className="p-4 my-2 shadow bg-white border">
				<h4 className="font-medium mb-2">Add Country</h4>
				<form onSubmit={submit}>
					<input
						type="text"
						className="block mb-3 rounded-sm"
						placeholder="Name"
						value={data.name}
						onChange={(e) => setData("name", e.target.value)}
					/>
					<input
						type="text"
						className="block mb-3 rounded-sm"
						placeholder="Country Code"
						maxLength={2}
						value={data.code}
						onChange={(e) => setData("code", e.target.value)}
					/>
					<button
						className="block bg-purple-600 text-white text-sm font-medium rounded-sm px-4 py-2"
						type="submit"
						disabled={processing}
					>
						Submit
					</button>
				</form>
			</div>
			<div className="py-6 text-lg">
				<h4 className="font-medium">All Countries</h4>
				<ul>
					{props.countries.map((c) => (
						<li key={c.Name}>
							{c.Flag} {c.Name}
						</li>
					))}
				</ul>
			</div>
		</Layout>
	);
}
