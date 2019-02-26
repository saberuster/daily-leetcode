import unittest
import twosum


class TestSolution(unittest.TestCase):
    def test_twosum(self):
        s = twosum.Solution()
        result = s.twoSum([2, 7, 11, 15], 9)
        self.assertEqual([0, 1], result)
