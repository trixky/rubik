type Axis = 'x' | 'y' | 'z';
type AxisList = Array<'x' | 'y' | 'z'>;

export default class Move {
	axe: Axis;
	otherAxis: AxisList;
	depth: -1 | 0 | 1;
	reverse: boolean;

	constructor(axe: Axis, depth: -1 | 0 | 1, reverse: boolean) {
		this.axe = axe;
		this.depth = depth;
		this.otherAxis = (['x', 'y', 'z'] as AxisList).filter((_axe) => _axe != axe);
		this.reverse = reverse;
	}
}
