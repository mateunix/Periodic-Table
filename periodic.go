package main

import (
	"fmt"
)

type chem struct {
	atomnum int64
	name    string
	symbol  string
	weight  float64
}

func printElementInfo(e chem) {
	fmt.Println("\nElement Information:\n")
	fmt.Printf("Name:    %s\n", e.name)
	fmt.Printf("Symbol:  %s\n", e.symbol)
	fmt.Printf("Atomic Number: %d\n", e.atomnum)
	fmt.Printf("Atomic Weight: %.5f\n", e.weight)

}

func main() {

	// User input for element symbol
	var input string

	fmt.Println("Type the element symbol(case-sensitive), for more information: ")
	fmt.Scanf("%s", &input)

	switch input {

	case "H":
		element := chem{1, "Hydrogen", "H", 1.00784}
		printElementInfo(element)
		fmt.Println("Description: Hydrogen is the lightest element and a key building block of the universe. It is essential for stars' fusion processes.")

	case "He":
		element := chem{2, "Helium", "He", 4.0026}
		printElementInfo(element)
		fmt.Println("Description: Helium is a noble gas, the second lightest element. It is inert and non-reactive at room temperature.")

	case "Li":
		element := chem{3, "Lithium", "Li", 6.94}
		printElementInfo(element)
		fmt.Println("Description: Lithium is a soft, silvery-white alkali metal. It is the lightest solid element and used in rechargeable batteries.")

	case "Be":
		element := chem{4, "Beryllium", "Be", 9.0122}
		printElementInfo(element)
		fmt.Println("Description: Beryllium is a strong, lightweight metal used in aerospace applications. It is toxic in dust form.")

	case "B":
		element := chem{5, "Boron", "B", 10.81}
		printElementInfo(element)
		fmt.Println("Description: Boron is a metalloid element essential for plant growth and found in borates. It has many industrial uses.")

	case "C":
		element := chem{6, "Carbon", "C", 12.011}
		printElementInfo(element)
		fmt.Println("Description: Carbon is the foundation of all life on Earth and the basis of organic chemistry. It has multiple allotropes.")

	case "N":
		element := chem{7, "Nitrogen", "N", 14.007}
		printElementInfo(element)
		fmt.Println("Description: Nitrogen makes up 78% of Earth's atmosphere and is a vital component of amino acids and nucleic acids.")

	case "O":
		element := chem{8, "Oxygen", "O", 15.999}
		printElementInfo(element)
		fmt.Println("Description: Oxygen is essential for respiration in most living organisms. It also supports combustion.")

	case "F":
		element := chem{9, "Fluorine", "F", 18.998}
		printElementInfo(element)
		fmt.Println("Description: Fluorine is a highly reactive halogen, the most electronegative element, and used in Teflon production.")

	case "Ne":
		element := chem{10, "Neon", "Ne", 20.180}
		printElementInfo(element)
		fmt.Println("Description: Neon is a noble gas that emits a bright red-orange glow when electrically excited. It is used in neon signs.")

	case "Na":
		element := chem{11, "Sodium", "Na", 22.990}
		printElementInfo(element)
		fmt.Println("Description: Sodium is a soft, highly reactive alkali metal. It forms a strong bond with chlorine to make table salt.")

	case "Mg":
		element := chem{12, "Magnesium", "Mg", 24.305}
		printElementInfo(element)
		fmt.Println("Description: Magnesium is a metal used in alloys and for biological processes. It is essential for plant and animal life.")

	case "Al":
		element := chem{13, "Aluminum", "Al", 26.982}
		printElementInfo(element)
		fmt.Println("Description: Aluminum is a lightweight metal used extensively in the aerospace, automotive, and packaging industries.")

	case "Si":
		element := chem{14, "Silicon", "Si", 28.085}
		printElementInfo(element)
		fmt.Println("Description: Silicon is a semiconductor material, fundamental in modern electronics, especially in computer chips.")

	case "P":
		element := chem{15, "Phosphorus", "P", 30.974}
		printElementInfo(element)
		fmt.Println("Description: Phosphorus is essential for life, forming DNA, RNA, and ATP. It exists in white, red, and black allotropes.")

	case "S":
		element := chem{16, "Sulfur", "S", 32.06}
		printElementInfo(element)
		fmt.Println("Description: Sulfur is a yellow non-metal that is found in nature and used in fertilizers, medicines, and industry.")

	case "Cl":
		element := chem{17, "Chlorine", "Cl", 35.45}
		printElementInfo(element)
		fmt.Println("Description: Chlorine is a halogen used in disinfectants and in the production of PVC and other chemicals.")

	case "Ar":
		element := chem{18, "Argon", "Ar", 39.948}
		printElementInfo(element)
		fmt.Println("Description: Argon is a colorless, odorless noble gas used in lighting and as an inert gas in industrial applications.")

	case "K":
		element := chem{19, "Potassium", "K", 39.098}
		printElementInfo(element)
		fmt.Println("Description: Potassium is an alkali metal essential for nerve function and found in many minerals and foods.")

	case "Ca":
		element := chem{20, "Calcium", "Ca", 40.078}
		printElementInfo(element)
		fmt.Println("Description: Calcium is essential for bone health and is a key element in cement and construction materials.")

	case "Sc":
		element := chem{21, "Scandium", "Sc", 44.956}
		printElementInfo(element)
		fmt.Println("Description: Scandium is a rare metal used in aluminum alloys, improving strength and resistance to corrosion.")

	case "Ti":
		element := chem{22, "Titanium", "Ti", 47.867}
		printElementInfo(element)
		fmt.Println("Description: Titanium is a strong, corrosion-resistant metal used in aerospace and medical implants.")

	case "V":
		element := chem{23, "Vanadium", "V", 50.942}
		printElementInfo(element)
		fmt.Println("Description: Vanadium is used in steel alloys to increase strength and durability, particularly in aerospace applications.")

	case "Cr":
		element := chem{24, "Chromium", "Cr", 52.00}
		printElementInfo(element)
		fmt.Println("Description: Chromium is used in stainless steel and chrome plating. It also has applications in pigments and metallurgy.")

	case "Mn":
		element := chem{25, "Manganese", "Mn", 54.938}
		printElementInfo(element)
		fmt.Println("Description: Manganese is essential for steel production and plays a role in human metabolism and bone health.")

	case "Fe":
		element := chem{26, "Iron", "Fe", 55.845}
		printElementInfo(element)
		fmt.Println("Description: Iron is a crucial element for life, particularly in oxygen transport, and is the most commonly used metal.")

	case "Ni":
		element := chem{27, "Nickel", "Ni", 58.693}
		printElementInfo(element)
		fmt.Println("Description: Nickel is used in stainless steel and batteries, and is also important for catalysis in chemical reactions.")

	case "Co":
		element := chem{28, "Cobalt", "Co", 58.933}
		printElementInfo(element)
		fmt.Println("Description: Cobalt is used in alloys for magnets and batteries, and is also essential for vitamin B12 synthesis.")

	case "Cu":
		element := chem{29, "Copper", "Cu", 63.546}
		printElementInfo(element)
		fmt.Println("Description: Copper is a highly conductive metal used in electrical wiring, plumbing, and coinage.")

	case "Zn":
		element := chem{30, "Zinc", "Zn", 65.38}
		printElementInfo(element)
		fmt.Println("Description: Zinc is used in galvanizing metals to prevent corrosion and is essential for immune function in humans.")

	case "Ga":
		element := chem{31, "Gallium", "Ga", 69.723}
		printElementInfo(element)
		fmt.Println("Description: Gallium is used in semiconductors and solar panels, and can melt just above room temperature.")

	case "Ge":
		element := chem{32, "Germanium", "Ge", 72.63}
		printElementInfo(element)
		fmt.Println("Description: Germanium is a semiconductor material used in transistors and fiber optics.")

	case "As":
		element := chem{33, "Arsenic", "As", 74.922}
		printElementInfo(element)
		fmt.Println("Description: Arsenic is a toxic metalloid and used in semiconductors and pesticides.")

	case "Se":
		element := chem{34, "Selenium", "Se", 78.971}
		printElementInfo(element)
		fmt.Println("Description: Selenium is essential for antioxidant activity in the body and is used in electronics and solar cells.")

	case "Br":
		element := chem{35, "Bromine", "Br", 79.904}
		printElementInfo(element)
		fmt.Println("Description: Bromine is a halogen used in flame retardants and water purification.")

	case "Kr":
		element := chem{36, "Krypton", "Kr", 83.798}
		printElementInfo(element)
		fmt.Println("Description: Krypton is a rare noble gas used in fluorescent lamps and in the production of certain lasers.")

	case "Rb":
		element := chem{37, "Rubidium", "Rb", 85.468}
		printElementInfo(element)
		fmt.Println("Description: Rubidium is an alkali metal used in research and in the production of specialty glasses.")

	case "Sr":
		element := chem{38, "Strontium", "Sr", 87.62}
		printElementInfo(element)
		fmt.Println("Description: Strontium is used in fireworks to produce red colors and in magnets and nuclear reactors.")

	case "Y":
		element := chem{39, "Yttrium", "Y", 88.906}
		printElementInfo(element)
		fmt.Println("Description: Yttrium is used in superconductors, LEDs, and in medical imaging.")

	case "Zr":
		element := chem{40, "Zirconium", "Zr", 91.224}
		printElementInfo(element)
		fmt.Println("Description: Zirconium is a strong, corrosion-resistant metal used in nuclear reactors and ceramics.")

	case "Nb":
		element := chem{41, "Niobium", "Nb", 92.906}
		printElementInfo(element)
		fmt.Println("Description: Niobium is used in superconducting magnets and in the production of stainless steel.")

	case "Mo":
		element := chem{42, "Molybdenum", "Mo", 95.95}
		printElementInfo(element)
		fmt.Println("Description: Molybdenum is a key component in steel production and is used in electrical contacts and filaments.")

	case "Tc":
		element := chem{43, "Technetium", "Tc", 98}
		printElementInfo(element)
		fmt.Println("Description: Technetium is a radioactive element used in medical imaging and is the first artificial element.")

	case "Ru":
		element := chem{44, "Ruthenium", "Ru", 101.07}
		printElementInfo(element)
		fmt.Println("Description: Ruthenium is used in electronics, as a catalyst, and in the production of platinum alloys.")

	case "Rh":
		element := chem{45, "Rhodium", "Rh", 102.91}
		printElementInfo(element)
		fmt.Println("Description: Rhodium is a rare, highly reflective metal used in catalytic converters and jewelry.")

	case "Pd":
		element := chem{46, "Palladium", "Pd", 106.42}
		printElementInfo(element)
		fmt.Println("Description: Palladium is used in catalytic converters, hydrogen storage, and in electronics.")

	case "Ag":
		element := chem{47, "Silver", "Ag", 107.87}
		printElementInfo(element)
		fmt.Println("Description: Silver is a precious metal known for its conductivity and is used in coins, jewelry, and mirrors.")

	case "Cd":
		element := chem{48, "Cadmium", "Cd", 112.41}
		printElementInfo(element)
		fmt.Println("Description: Cadmium is toxic and is used in batteries, pigments, and as a corrosion-resistant coating.")

	case "In":
		element := chem{49, "Indium", "In", 114.82}
		printElementInfo(element)
		fmt.Println("Description: Indium is used in electronics, particularly in touchscreens and flat-panel displays.")

	case "Sn":
		element := chem{50, "Tin", "Sn", 118.71}
		printElementInfo(element)
		fmt.Println("Description: Tin is used in soldering and as a protective coating for other metals to prevent corrosion.")

	case "Sb":
		element := chem{51, "Antimony", "Sb", 121.76}
		printElementInfo(element)
		fmt.Println("Description: Antimony is used in flame retardants and alloys, and has been known since ancient times.")

	case "I":
		element := chem{52, "Iodine", "I", 126.90}
		printElementInfo(element)
		fmt.Println("Description: Iodine is essential for thyroid hormone production and used in antiseptics and pharmaceuticals.")

	case "Te":
		element := chem{53, "Tellurium", "Te", 127.60}
		printElementInfo(element)
		fmt.Println("Description: Tellurium is used in alloys, solar panels, and thermoelectric devices.")

	case "Xe":
		element := chem{54, "Xenon", "Xe", 131.29}
		printElementInfo(element)
		fmt.Println("Description: Xenon is a noble gas used in high-intensity lamps, anesthesia, and ion propulsion systems.")

	case "Cs":
		element := chem{55, "Cesium", "Cs", 132.91}
		printElementInfo(element)
		fmt.Println("Description: Cesium is used in atomic clocks and in oil drilling applications.")

	case "Ba":
		element := chem{56, "Barium", "Ba", 137.33}
		printElementInfo(element)
		fmt.Println("Description: Barium is used in oil drilling, X-ray imaging, and in fireworks.")

	case "La":
		element := chem{57, "Lanthanum", "La", 138.91}
		printElementInfo(element)
		fmt.Println("Description: Lanthanum is used in catalysts, lighting, and as a component in various rare-earth alloys.")

	case "Ce":
		element := chem{58, "Cerium", "Ce", 140.12}
		printElementInfo(element)
		fmt.Println("Description: Cerium is a rare earth metal used in catalytic converters and in glass polishing.")

	case "Pr":
		element := chem{59, "Praseodymium", "Pr", 140.91}
		printElementInfo(element)
		fmt.Println("Description: Praseodymium is used in alloys, magnets, and in the production of green dyes.")

	case "Nd":
		element := chem{60, "Neodymium", "Nd", 144.24}
		printElementInfo(element)
		fmt.Println("Description: Neodymium is used in high-strength magnets, lasers, and in electronics.")

	case "Pm":
		element := chem{61, "Promethium", "Pm", 145}
		printElementInfo(element)
		fmt.Println("Description: Promethium is a rare, radioactive metal used in atomic batteries and as a source of radiation.")

	case "Sm":
		element := chem{62, "Samarium", "Sm", 150.36}
		printElementInfo(element)
		fmt.Println("Description: Samarium is used in magnets, neutron capture therapy, and in carbon arc lights.")
	case "Eu":

		element := chem{63, "Europium", "Eu", 151.98}
		printElementInfo(element)
		fmt.Println("Description: Europium is used in phosphorescent and fluorescent applications, especially in LED displays.")

	case "Gd":
		element := chem{64, "Gadolinium", "Gd", 157.25}
		printElementInfo(element)
		fmt.Println("Description: Gadolinium is used in MRI contrast agents and neutron radiography.")

	case "Tb":
		element := chem{65, "Terbium", "Tb", 158.93}
		printElementInfo(element)
		fmt.Println("Description: Terbium is used in solid-state devices, phosphors, and in high-efficiency fluorescent lighting.")

	case "Dy":
		element := chem{66, "Dysprosium", "Dy", 162.50}
		printElementInfo(element)
		fmt.Println("Description: Dysprosium is used in magnets and is essential for data storage and electric vehicle motors.")

	case "Ho":
		element := chem{67, "Holmium", "Ho", 164.93}
		printElementInfo(element)
		fmt.Println("Description: Holmium is used in magnets, lasers, and in various chemical applications.")

	case "Er":
		element := chem{68, "Erbium", "Er", 167.26}
		printElementInfo(element)
		fmt.Println("Description: Erbium is used in optical fiber amplifiers and in medical lasers for surgery.")

	case "Tm":
		element := chem{69, "Thulium", "Tm", 168.93}
		printElementInfo(element)
		fmt.Println("Description: Thulium is used in X-ray machines and in lasers for medical treatments.")

	case "Yb":
		element := chem{70, "Ytterbium", "Yb", 173.04}
		printElementInfo(element)
		fmt.Println("Description: Ytterbium is used in fiber-optic communications and as a dopant in lasers.")

	case "Lu":
		element := chem{71, "Lutetium", "Lu", 175.00}
		printElementInfo(element)
		fmt.Println("Description: Lutetium is used in PET scanners and in catalysis for petroleum refining.")

	case "Hf":
		element := chem{72, "Hafnium", "Hf", 178.49}
		printElementInfo(element)
		fmt.Println("Description: Hafnium is used in nuclear reactors, in high-temperature alloys, and as a semiconductor material.")

	case "Ta":
		element := chem{73, "Tantalum", "Ta", 180.95}
		printElementInfo(element)
		fmt.Println("Description: Tantalum is highly corrosion-resistant and is used in electronics and in medical implants.")

	case "W":
		element := chem{74, "Tungsten", "W", 183.84}
		printElementInfo(element)
		fmt.Println("Description: Tungsten has the highest melting point of any metal and is used in lightbulb filaments and in aerospace.")

	case "Re":
		element := chem{75, "Rhenium", "Re", 186.21}
		printElementInfo(element)
		fmt.Println("Description: Rhenium is used in high-temperature superalloys and in the aerospace industry.")

	case "Os":
		element := chem{76, "Osmium", "Os", 190.23}
		printElementInfo(element)
		fmt.Println("Description: Osmium is the densest naturally occurring element, used in fountain pen tips and electrical contacts.")

	case "Ir":
		element := chem{77, "Iridium", "Ir", 192.22}
		printElementInfo(element)
		fmt.Println("Description: Iridium is highly resistant to corrosion and is used in spark plugs and in aerospace applications.")

	case "Pt":
		element := chem{78, "Platinum", "Pt", 195.08}
		printElementInfo(element)
		fmt.Println("Description: Platinum is a precious metal used in catalytic converters, jewelry, and various industrial processes.")

	case "Au":
		element := chem{79, "Gold", "Au", 196.97}
		printElementInfo(element)
		fmt.Println("Description: Gold is a soft, yellow metal used in jewelry, electronics, and as a store of value.")

	case "Hg":
		element := chem{80, "Mercury", "Hg", 200.59}
		printElementInfo(element)
		fmt.Println("Description: Mercury is the only metal that is liquid at room temperature and is used in thermometers and switches.")

	case "Tl":
		element := chem{81, "Thallium", "Tl", 204.38}
		printElementInfo(element)
		fmt.Println("Description: Thallium is a toxic metal used in semiconductors and in the production of optical materials.")

	case "Pb":
		element := chem{82, "Lead", "Pb", 207.2}
		printElementInfo(element)
		fmt.Println("Description: Lead is a toxic metal historically used in pipes, paint, and batteries, but now largely phased out.")

	case "Bi":
		element := chem{83, "Bismuth", "Bi", 208.98}
		printElementInfo(element)
		fmt.Println("Description: Bismuth is a brittle, non-toxic metal used in cosmetics, pharmaceuticals, and low-melting alloys.")

	case "Po":
		element := chem{84, "Polonium", "Po", 209}
		printElementInfo(element)
		fmt.Println("Description: Polonium is a highly radioactive element used in anti-static devices and has been associated with poisoning.")

	case "At":
		element := chem{85, "Astatine", "At", 210}
		printElementInfo(element)
		fmt.Println("Description: Astatine is a rare, radioactive halogen with no significant commercial uses due to its instability.")

	case "Rn":
		element := chem{86, "Radon", "Rn", 222}
		printElementInfo(element)
		fmt.Println("Description: Radon is a colorless, odorless radioactive gas that occurs naturally from the decay of uranium.")

	case "Fr":
		element := chem{87, "Francium", "Fr", 223}
		printElementInfo(element)
		fmt.Println("Description: Francium is a highly radioactive alkali metal and the least stable of the first 101 elements.")

	case "Ra":
		element := chem{88, "Radium", "Ra", 226}
		printElementInfo(element)
		fmt.Println("Description: Radium is a highly radioactive metal once used in luminous paints but now mostly replaced due to safety concerns.")

	case "Ac":
		element := chem{89, "Actinium", "Ac", 227}
		printElementInfo(element)
		fmt.Println("Description: Actinium is a radioactive metal used in radiation therapy and as a neutron source.")

	case "Th":
		element := chem{90, "Thorium", "Th", 232.04}
		printElementInfo(element)
		fmt.Println("Description: Thorium is a radioactive metal being researched for use in nuclear reactors and as a potential alternative fuel.")

	case "Pa":
		element := chem{91, "Protactinium", "Pa", 231.04}
		printElementInfo(element)
		fmt.Println("Description: Protactinium is a radioactive element used in nuclear research but has limited practical applications.")

	case "U":
		element := chem{92, "Uranium", "U", 238.03}
		printElementInfo(element)
		fmt.Println("Description: Uranium is a radioactive metal used as fuel in nuclear reactors and in the production of nuclear weapons.")

	case "Np":
		element := chem{93, "Neptunium", "Np", 237}
		printElementInfo(element)
		fmt.Println("Description: Neptunium is a radioactive actinide metal and was the first transuranic element to be synthesized.")

	case "Pu":
		element := chem{94, "Plutonium", "Pu", 244}
		printElementInfo(element)
		fmt.Println("Description: Plutonium is a radioactive metal used in nuclear reactors and in the manufacture of nuclear weapons.")

	case "Am":
		element := chem{95, "Americium", "Am", 243}
		printElementInfo(element)
		fmt.Println("Description: Americium is used in smoke detectors and in some types of radiation therapy.")

	case "Cm":
		element := chem{96, "Curium", "Cm", 247}
		printElementInfo(element)
		fmt.Println("Description: Curium is a synthetic radioactive element used in radiation therapy and in scientific research.")

	case "Bk":
		element := chem{97, "Berkelium", "Bk", 247}
		printElementInfo(element)
		fmt.Println("Description: Berkelium is a synthetic element named after Berkeley, California, and is used in scientific research.")

	case "Cf":
		element := chem{98, "Californium", "Cf", 251}
		printElementInfo(element)
		fmt.Println("Description: Californium is a synthetic radioactive element used in nuclear reactors and as a neutron source.")

	case "Es":
		element := chem{99, "Einsteinium", "Es", 252}
		printElementInfo(element)
		fmt.Println("Description: Einsteinium is a synthetic element named after Albert Einstein and is used in research and the production of heavier elements.")

	case "Fm":
		element := chem{100, "Fermium", "Fm", 257}
		printElementInfo(element)
		fmt.Println("Description: Fermium is a synthetic element used in scientific research, especially in nuclear chemistry and particle physics.")

	case "Md":
		element := chem{101, "Mendelevium", "Md", 258}
		printElementInfo(element)
		fmt.Println("Description: Mendelevium is a synthetic element named after Dmitri Mendeleev, used in research involving nuclear reactions.")

	case "No":
		element := chem{102, "Nobelium", "No", 259}
		printElementInfo(element)
		fmt.Println("Description: Nobelium is a synthetic element named after Alfred Nobel, used in particle physics experiments.")

	case "Lr":
		element := chem{103, "Lawrencium", "Lr", 262}
		printElementInfo(element)
		fmt.Println("Description: Lawrencium is a synthetic element used in nuclear physics research, named after Ernest O. Lawrence.")

	case "Rf":
		element := chem{104, "Rutherfordium", "Rf", 267}
		printElementInfo(element)
		fmt.Println("Description: Rutherfordium is a synthetic element used in nuclear research, named after physicist Ernest Rutherford.")

	case "Db":
		element := chem{105, "Dubnium", "Db", 270}
		printElementInfo(element)
		fmt.Println("Description: Dubnium is a synthetic element used in nuclear physics experiments.")

	case "Sg":
		element := chem{106, "Seaborgium", "Sg", 271}
		printElementInfo(element)
		fmt.Println("Description: Seaborgium is a synthetic element named after Glenn T. Seaborg, used in nuclear research.")

	case "Bh":
		element := chem{107, "Bohrium", "Bh", 270}
		printElementInfo(element)
		fmt.Println("Description: Bohrium is a synthetic element used in nuclear research, named after Niels Bohr.")

	case "Hs":
		element := chem{108, "Hassium", "Hs", 277}
		printElementInfo(element)
		fmt.Println("Description: Hassium is a synthetic element used in nuclear research, named after the German state of Hesse.")

	case "Mt":
		element := chem{109, "Meitnerium", "Mt", 278}
		printElementInfo(element)
		fmt.Println("Description: Meitnerium is a synthetic element named after Lise Meitner, used in nuclear research.")

	case "Ds":
		element := chem{110, "Darmstadtium", "Ds", 281}
		printElementInfo(element)
		fmt.Println("Description: Darmstadtium is a synthetic element named after Darmstadt, Germany, used in nuclear physics.")

	case "Rg":
		element := chem{111, "Roentgenium", "Rg", 280}
		printElementInfo(element)
		fmt.Println("Description: Roentgenium is a synthetic element named after Wilhelm Röntgen, used in nuclear research.")

	case "Cn":
		element := chem{112, "Copernicium", "Cn", 285}
		printElementInfo(element)
		fmt.Println("Description: Copernicium is a synthetic element named after Nicolaus Copernicus, used in particle physics.")

	case "Nh":
		element := chem{113, "Nihonium", "Nh", 284}
		printElementInfo(element)
		fmt.Println("Description: Nihonium is a synthetic element named after Japan, used in nuclear research.")

	case "Fl":
		element := chem{114, "Flerovium", "Fl", 289}
		printElementInfo(element)
		fmt.Println("Description: Flerovium is a synthetic element named after the Flerov Laboratory of Nuclear Reactions.")

	case "Mc":
		element := chem{115, "Moscovium", "Mc", 288}
		printElementInfo(element)
		fmt.Println("Description: Moscovium is a synthetic element named after Moscow, used in nuclear physics.")

	case "Lv":
		element := chem{116, "Livermorium", "Lv", 293}
		printElementInfo(element)
		fmt.Println("Description: Livermorium is a synthetic element named after the Lawrence Livermore National Laboratory.")

	case "Ts":
		element := chem{117, "Tennessine", "Ts", 294}
		printElementInfo(element)
		fmt.Println("Description: Tennessine is a synthetic element named after Tennessee, used in nuclear research.")

	case "Og":
		element := chem{118, "Oganesson", "Og", 294}
		printElementInfo(element)
		fmt.Println("Description: Oganesson is a synthetic element named after Yuri Oganessian, used in particle physics.")

	default:
		fmt.Println("Element not found!")
	}
}
