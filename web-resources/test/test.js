import { assert } from 'chai';

console.log("Starting unit tests");

describe('Array', function() {
  describe('#indexOf()', function() {
    it('should return -1 when the value is not present', function() {
    	console.log("First unit test")
      assert.equal([1, 2, 3].indexOf(4), -1);
    });
  });
});
