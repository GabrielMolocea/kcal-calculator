import { useEffect, useState } from "react";
import CaloriesForm from "./components/form/caloriesForm";

function App() {
	const [targetCalories, setTargetCalories] = useState(0);
	const [totalKcal, setTotalKcal] = useState(0);
	const [foodList, setFoodList] = useState<
		{ id: number; name: string; calories: number }[]
	>([]);

	useEffect(() => {
		getTargetCalories();
		getFoodList();
	}, []);

	async function getTargetCalories() {
		await fetch("http://localhost:8080/targetkcal")
			.then((response) => response.json())
			.then((data) => setTargetCalories(data.targetKcal));
	}

	async function getFoodList() {
		const response = await fetch("http://localhost:8080/foods", {
			method: "GET",
			headers: {
				"Content-Type": "application/json",
			},
		});
		setFoodList(await response.json());
	}

	async function resetTargetCalories() {
		await fetch("http://localhost:8080/resttargetkcal", {
			method: "POST",
			headers: {
				"Content-Type": "application/json",
			},
		})
			.then((res) => res.json())
			.then((data) => setTargetCalories(data.targetKcal))
			.then(() => getTargetCalories())
			.then(() => setTotalKcal(0));
	}

	return (
		<div className="min-h-screen bg-gray-100 p-4 flex flex-col justify-center items-center">
			<div className="max-w-2xl mx-auto bg-white p-6 rounded-lg shadow-md">
				<h1 className="text-3xl font-bold mb-4">Calorie Calculator</h1>
				<h2 className="text-xl mb-4">Target Calories: {targetCalories}</h2>
				<CaloriesForm setTotalKcal={setTotalKcal} foodList={foodList} getTargetCalories={getTargetCalories} />
				<h2 className="text-xl mt-4">
					You have eaten this much calories today: {totalKcal}
				</h2>
				<button
					onClick={resetTargetCalories}
					className="mt-4 bg-blue-500 text-white py-2 px-4 rounded-lg hover:bg-blue-600"
				>
					Reset target
				</button>
			</div>
		</div>
	);
}

export default App;
