from aoc2019.day_six.planetary_iterator import BFSPlanetaryIterator


def get_depth(tree, name):
    """
    Simple utility method that iterates through components in a PlanetaryTree.
    It determines the minimum distance from root to named node.
    """
    for node in tree:
        if node[0].name == name:
            return node[1]  # "depth" it was found at

    # Element not found
    raise ValueError


def get_min_dist(tree, name_one, name_two):
    """
    Utility function that returns the minimum distance between components with
    passed names.
    """
    tree_iter = iter(tree)
    tree_iter.__next__()  # consume level 0 component
    component = tree_iter.__next__()  # get first level 1 component
    # Iterate through children with depth = 1
    try:
        while component[1] < 2:
            try:
                sub_tree = PlanetaryTree(component[0])
                # child sub-trees will have smaller edge distances
                return get_min_dist(sub_tree, name_one, name_two)
            except ValueError:
                # sub tree does not contain both components - do nothing
                pass

            # iterate to next component
            component = tree_iter.__next__()
    except StopIteration:
        # no more elements - do nothing
        pass

    # passed tree has lowest level shared root for named components
    return get_depth(tree, name_one) + get_depth(tree, name_two) - 2


class SatelliteLeaf:
    """This a class for satellite entities - such as planets"""
    def __init__(self, name):
        self.name = name

    def count(self):
        """
        Simple utility method to return the number of nodes in this leaf,
         which is always 1.
        """
        return 1

    def count_orbits(self, level):
        """
        This method returns the number of orbiting entities - 1 as it
        is a leaf.
         """
        return level

    def __str__(self):
        return 'leaf: {}'.format(self.name)


class CentralMassComposite:
    """
    This class is a composite component for central mass entities - such as
    what planets orbit around.
    """
    def __init__(self, name):
        self.name = name
        self.satellites = []  # collection of orbiting entities

    def count(self):
        """
        Simple utility method to recursively return the number of nodes in
        this composite sub tree.
        """
        count = 1
        for satellite in self.satellites:
            count += satellite.count()
        return count

    def count_orbits(self, level):
        """
        This method returns recursively the number of all orbiting entities,
         both direct and indirect.
        """
        count = 0
        for satellite in self.satellites:
            count += satellite.count_orbits(level + 1)

        return level + count

    def __str__(self):
        return self.name


class PlanetaryTree:
    """
    This is a collection of planetary tree components - CentralMassComposites
    and SatelliteLeafs. It is an iterable collection.
    """

    def __init__(self, root):
        """Basic initialization method."""
        self.root = root

    def contains(self, key):
        """
        Utility method to return true if a component with passed name is
        present in tree
        """
        # Iterate through all components
        for component in self:
            if component[0].name == key:
                return True  # found component

        # No matching components found
        return False

    def count_all_orbits(self):
        """
        Utility method that returns the sum of all direct and indirect orbits.
        """
        return self.root.count_orbits(0)

    def __iter__(self):
        """Initializes and returns the iterator object"""
        return BFSPlanetaryIterator(self.root)
