using UnityEngine;
using System.Collections;

public class Bomb : MonoBehaviour {

    private int id_ = 0;
    private Vector3 scale_ = new Vector3(10, 10, 20);

	// Use this for initialization
	void Start () {
	
	}
	
	// Update is called once per frame
	void Update () {
	
	}

    public void Init(int id, int x, int y)
    {
        transform.localScale = scale_;
        transform.position = new Vector3(x, 0, y);
        id_ = id;
    }

    public void Blast()
    {
        GameObject o1 = GameObject.CreatePrimitive(PrimitiveType.Cube);
        o1.transform.position = transform.position;
        o1.transform.localScale = new Vector3(200, 10, 10);
        o1.AddComponent<Blast>();

        GameObject o2 = GameObject.CreatePrimitive(PrimitiveType.Cube);
        o2.AddComponent<Blast>();
        o2.transform.position = transform.position;
        o2.transform.localScale = new Vector3(10, 10, 200);

        Destroy(gameObject);
    }
}
