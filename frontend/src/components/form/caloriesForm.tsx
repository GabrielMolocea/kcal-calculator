import { useEffect, useState, useCallback } from "react";
import { debounce } from "../util/debouncer";
import "../../index.css";

interface CaloriesFormProps {
	setTotalKcal: (kcal: number) => void;
	foodList: { name: string; calories: number }[];
}

const CaloriesForm: React.FC<CaloriesFormProps> = ({
	setTotalKcal,
	foodList = [],
}) => {
	const [food, setFood] = useState({ name: "", calories: 0 });
	const [searchTerm, setSearchTerm] = useState("");
	const [filteredFoodList, setFilteredFoodList] = useState(foodList);

	useEffect(() => {
		const results = foodList.filter((food) =>
			food.name.toLowerCase().includes(searchTerm.toLowerCase())
		);
		setFilteredFoodList(results);
	}, [searchTerm, foodList]);

	const debouncedSearch = useCallback(
		debounce((value: string) => {
			setSearchTerm(value);
		}, 300),
		[]
	);

	async function addFood(event: any) {
		event.preventDefault();
		await fetch("http://localhost:8080/foods", {
			method: "POST",
			headers: {
				"Content-Type": "application/json",
			},
			body: JSON.stringify(food),
		});

		console.log(food);
	}

	const handleSearchChange = (event: React.ChangeEvent<HTMLInputElement>) => {
		debouncedSearch(event.target.value);
	};

	const handleFoodSelect = (selectedFood: {
		name: string;
		calories: number;
	}) => {
		setFood(selectedFood);
	};

	return (
		<>
			<div>Add Food</div>
			<input type="text" placeholder="Search food" onChange={handleSearchChange} />
			<input type="Number" placeholder="Search food" onChange={handleSearchChange} />
			<button type="submit" onClick={addFood}>
				Add Food
			</button>
		</>
	);
};

export default CaloriesForm;
