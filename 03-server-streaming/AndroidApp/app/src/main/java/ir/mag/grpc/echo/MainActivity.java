package ir.mag.grpc.echo;

import android.os.Bundle;
import android.util.Log;

import androidx.appcompat.app.AppCompatActivity;

import com.google.android.material.button.MaterialButton;

import java.util.Iterator;

import io.grpc.ManagedChannel;
import io.grpc.ManagedChannelBuilder;

public class MainActivity extends AppCompatActivity {

	private static final String TAG = "MainActivity";

	@Override
	protected void onCreate(Bundle savedInstanceState) {
		super.onCreate(savedInstanceState);
		setContentView(R.layout.activity_main);

		String host = "192.168.1.46";
		int port = 8888;
		ManagedChannel channel = ManagedChannelBuilder.forAddress(host, port).usePlaintext().build();


		MaterialButton materialButton = findViewById(R.id.connect);
		materialButton.setOnClickListener(view -> {
			Log.d(TAG, "onCreate: clicked");
			Log.d(TAG, "globalCounter: in the global counter");
			CounterParam request = CounterParam.newBuilder().build();
			Iterator<GlobalCounterResponse> features;
			Log.d(TAG, "globalCounter: set features");
			features = EchoGrpc.newBlockingStub(channel).globalCounter(request);

			Log.d(TAG, "globalCounter: starting the while loop");
			while (features.hasNext()) {
				Log.d(TAG, "globalCounter: waiting");
				GlobalCounterResponse feature = features.next();
				Log.d(TAG, "globalCounter: " + feature.getCounter());
			}
		});

	}

}